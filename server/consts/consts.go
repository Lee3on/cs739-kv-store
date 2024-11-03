package consts

import "time"

const (
	Success       = 0
	InternalError = -1
	KeyNotFound   = 1
	Redirect      = 2
)

const (
	RaftPortBase = 5000
)

const (
	KVStoreCapacity    = 20
	TTL                = 10
	KVStoreEvictionTTL = 10 * time.Second
)

const (
	RaftServerListFileName = "./config/raft_server_list"
	KVServerListFileName   = "./config/kv_server_list"
)
