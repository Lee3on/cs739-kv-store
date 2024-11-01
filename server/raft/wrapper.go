package raft

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"go.etcd.io/etcd/client/pkg/v3/types"
	"go.etcd.io/etcd/raft/v3"
	"go.etcd.io/etcd/raft/v3/raftpb"
	"go.etcd.io/etcd/server/v3/etcdserver/api/rafthttp"
	stats "go.etcd.io/etcd/server/v3/etcdserver/api/v2stats"
	"go.etcd.io/etcd/server/v3/wal"
	"go.etcd.io/etcd/server/v3/wal/walpb"
	"go.uber.org/zap"
)

type Wrapper struct {
	id    uint64            // Unique ID for the node
	peers map[uint64]string // Peer node IDs and their addresses

	node      raft.Node           // Raft Node representing the consensus instance
	wal       *wal.WAL            // Write-Ahead Log for persisting entries
	storage   *raft.MemoryStorage // Memory storage to hold entries and snapshots in memory
	transport *rafthttp.Transport // Network transport for Raft messages between nodes
	ticker    *time.Ticker        // Ticker for Raft election and heartbeat timing

	proposeC     <-chan string   // Channel for proposed data to be committed
	commitC      chan<- []string // Channel to send committed entries
	appliedIndex uint64          // Index of the last applied entry
	stopc        chan struct{}
}

// Constructor for Wrapper, initializes Raft node, WAL, storage, and transport
func NewWrapper(id uint64, peers map[uint64]string, proposeC <-chan string, commitC chan<- []string) (*Wrapper, error) {
	wrapper := &Wrapper{
		id:    id,
		peers: peers,
		stopc: make(chan struct{}),
	}

	// Initialize Raft Peers
	npeers := make([]raft.Peer, 0, len(peers))
	for i := range peers {
		npeers = append(npeers, raft.Peer{ID: i})
	}

	// Initialize in-memory Raft storage
	storage := raft.NewMemoryStorage()
	waldir := fmt.Sprintf("./storage/wal/%d", id) // Directory for WAL files
	oldwal := wal.Exist(waldir)
	if !wal.Exist(waldir) {
		// Create WAL directory if it doesn't exist
		err := os.Mkdir(waldir, 0705)
		if err != nil {
			return nil, err
		}
		// Create a new WAL
		wal0, err := wal.Create(zap.NewExample(), waldir, nil)
		if err != nil {
			return nil, err
		}
		wal0.Close()
	}
	// Open existing WAL, or the newly created WAL
	wal0, err := wal.Open(zap.NewExample(), waldir, walpb.Snapshot{})
	if err != nil {
		return nil, err
	}
	_, hardState, entries, err := wal0.ReadAll()
	if err != nil {
		return nil, err
	}
	// Restore hard state and entries to storage
	storage.SetHardState(hardState)
	storage.Append(entries)

	// Configuration for the Raft node
	config := &raft.Config{
		ID:              id,
		ElectionTick:    10,
		HeartbeatTick:   1,
		Storage:         storage,
		MaxSizePerMsg:   4096,
		MaxInflightMsgs: 9192,
	}
	// Start or restart Raft node based on existence of old WAL
	if oldwal {
		wrapper.node = raft.RestartNode(config)
	} else {
		wrapper.node = raft.StartNode(config, npeers)
	}

	// Setup transport for communication between Raft nodes
	transport := &rafthttp.Transport{
		Logger:      zap.NewExample(),
		ID:          types.ID(id),
		ClusterID:   0x1000,
		Raft:        wrapper,
		ServerStats: stats.NewServerStats("", ""),
		LeaderStats: stats.NewLeaderStats(zap.NewExample(), strconv.Itoa(int(id))),
		ErrorC:      make(chan error),
	}
	transport.Start()
	for i, addr := range peers {
		if i != id {
			transport.AddPeer(types.ID(i), []string{addr})
		}
	}

	// Assign initialized components to wrapper struct
	wrapper.wal = wal0
	wrapper.storage = storage
	wrapper.transport = transport
	wrapper.ticker = time.NewTicker(1 * time.Second)
	wrapper.proposeC = proposeC
	wrapper.commitC = commitC

	// Start goroutines for handling proposals, transport, and Raft operations
	go wrapper.RunPropose()
	go wrapper.RunTransport()
	go wrapper.Run()

	return wrapper, nil
}

// RunPropose listens on proposeC and forwards proposals to the Raft node
func (w *Wrapper) RunPropose() {
	for {
		select {
		case data := <-w.proposeC:
			w.node.Propose(context.TODO(), []byte(data))
		case <-w.stopc:
			return
		}
	}
}

// RunTransport listens for HTTP requests to serve Raft messages to peers
func (w *Wrapper) RunTransport() {
	url, _ := url.Parse(w.peers[w.id])
	server := &http.Server{
		Addr:    url.Host,
		Handler: w.transport.Handler(),
	}
	go func() {
		<-w.stopc
		server.Close()
	}()
	server.ListenAndServe()
}

// Run handles the Raft node's periodic operations, including tick, Ready handling, and advancing state
func (w *Wrapper) Run() {
	for {
		select {
		case <-w.ticker.C:
			// Ticking advances Raft time (handles election and heartbeat)
			w.node.Tick()
		case rd := <-w.node.Ready():
			// Save WAL and memory storage entries
			w.wal.Save(rd.HardState, rd.Entries)
			w.storage.Append(rd.Entries)
			w.transport.Send(rd.Messages)

			// Process committed entries (log entries that have been agreed upon)
			entries := rd.CommittedEntries
			if len(entries) > 0 {
				entries = entries[w.appliedIndex+1-entries[0].Index:]
				datas := make([]string, 0, len(entries))
				for _, entry := range entries {
					switch entry.Type {
					case raftpb.EntryNormal:
						// Process normal log entries with actual data
						if len(entry.Data) == 0 {
							break
						}
						datas = append(datas, string(entry.Data))
					case raftpb.EntryConfChange:
						// Handle configuration changes
						var cc raftpb.ConfChange
						cc.Unmarshal(entry.Data)
						w.node.ApplyConfChange(cc)
						switch cc.Type {
						case raftpb.ConfChangeAddNode:
							// Add new peer to transport
							if len(cc.Context) > 0 {
								w.transport.AddPeer(types.ID(cc.NodeID), []string{string(cc.Context)})
							}
						case raftpb.ConfChangeRemoveNode:
							// Remove peer from transport
							w.transport.RemovePeer(types.ID(cc.NodeID))
							if cc.NodeID == w.id {
								// Our node has been removed from the cluster; shutdown gracefully
								log.Printf("Node %d has been removed from the cluster. Shutting down.\n", w.id)
								w.Stop()
								return
							}
						}
					}
				}
				w.commitC <- datas                             // Commit the processed data
				w.appliedIndex = entries[len(entries)-1].Index // Update applied index
			}

			w.node.Advance() // Signal that Ready entries are processed
		}
	}
}

func (w *Wrapper) Shutdown() {
	log.Printf("Node %d is shutting down to simulate failure.\n", w.id)
	w.ReportUnreachable(w.id)
	w.Stop()
}

func (w *Wrapper) Start() error {
	cc := raftpb.ConfChange{
		Type:    raftpb.ConfChangeAddNode,
		NodeID:  w.id,
		Context: []byte(w.peers[w.id]),
	}
	return w.node.ProposeConfChange(context.TODO(), cc)
}

func (w *Wrapper) Remove() error {
	cc := raftpb.ConfChange{
		Type:   raftpb.ConfChangeRemoveNode,
		NodeID: w.id,
	}
	return w.node.ProposeConfChange(context.TODO(), cc)
}

func (w *Wrapper) Stop() {
	w.ticker.Stop()
	w.node.Stop()
	w.transport.Stop()
	w.wal.Close()
	log.Printf("Node %d has stopped.\n", w.id)
}

func (w *Wrapper) IsLeader() bool {
	return w.node.Status().Lead == w.id
}

func (w *Wrapper) GetLeader() uint64 {
	return w.node.Status().Lead
}

// Process implements Raft transport interface to handle messages from other nodes
func (w *Wrapper) Process(ctx context.Context, m raftpb.Message) error {
	return w.node.Step(ctx, m)
}

// IsIDRemoved checks if a node has been removed
func (w *Wrapper) IsIDRemoved(id uint64) bool {
	return false
}

// ReportUnreachable reports a node as unreachable in the Raft protocol
func (w *Wrapper) ReportUnreachable(id uint64) {
	w.node.ReportUnreachable(id)
}

// ReportSnapshot reports the status of a snapshot transfer in Raft
func (w *Wrapper) ReportSnapshot(id uint64, status raft.SnapshotStatus) {
	// No-op for now
}
