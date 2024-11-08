package raft

import (
	"reflect"
	"testing"

	"go.etcd.io/etcd/raft/v3/raftpb"
)

func TestProcessMessages(t *testing.T) {
	cases := []struct {
		name             string
		confState        raftpb.ConfState
		InputMessages    []raftpb.Message
		ExpectedMessages []raftpb.Message
	}{
		{
			name: "only one snapshot message",
			confState: raftpb.ConfState{
				Voters: []uint64{2, 6, 8, 10},
			},
			InputMessages: []raftpb.Message{
				{
					Type: raftpb.MsgSnap,
					To:   8,
					Snapshot: raftpb.Snapshot{
						Metadata: raftpb.SnapshotMetadata{
							Index: 100,
							Term:  3,
							ConfState: raftpb.ConfState{
								Voters:    []uint64{2, 6, 8},
								AutoLeave: true,
							},
						},
					},
				},
			},
			ExpectedMessages: []raftpb.Message{
				{
					Type: raftpb.MsgSnap,
					To:   8,
					Snapshot: raftpb.Snapshot{
						Metadata: raftpb.SnapshotMetadata{
							Index: 100,
							Term:  3,
							ConfState: raftpb.ConfState{
								Voters: []uint64{2, 6, 8, 10},
							},
						},
					},
				},
			},
		},
		{
			name: "one snapshot message and one other message",
			confState: raftpb.ConfState{
				Voters: []uint64{2, 7, 8, 12},
			},
			InputMessages: []raftpb.Message{
				{
					Type: raftpb.MsgSnap,
					To:   8,
					Snapshot: raftpb.Snapshot{
						Metadata: raftpb.SnapshotMetadata{
							Index: 100,
							Term:  3,
							ConfState: raftpb.ConfState{
								Voters:    []uint64{2, 6, 8},
								AutoLeave: true,
							},
						},
					},
				},
				{
					Type: raftpb.MsgApp,
					From: 6,
					To:   8,
				},
			},
			ExpectedMessages: []raftpb.Message{
				{
					Type: raftpb.MsgSnap,
					To:   8,
					Snapshot: raftpb.Snapshot{
						Metadata: raftpb.SnapshotMetadata{
							Index: 100,
							Term:  3,
							ConfState: raftpb.ConfState{
								Voters: []uint64{2, 7, 8, 12},
							},
						},
					},
				},
				{
					Type: raftpb.MsgApp,
					From: 6,
					To:   8,
				},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			rn := &RaftNode{
				confState: tc.confState,
			}

			outputMessages := rn.processMessages(tc.InputMessages)

			if !reflect.DeepEqual(outputMessages, tc.ExpectedMessages) {
				t.Fatalf("Unexpected messages, expected: %v, got %v", tc.ExpectedMessages, outputMessages)
			}
		})
	}
}
