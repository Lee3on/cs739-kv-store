package consts

const (
	Success       = 0
	InternalError = -1
	KeyNotFound   = 1
	Redirect      = 2
)

const (
	KVStoreCapacity = 20
)

const (
	RaftServerListFileName = "./config/raft_server_list"
	KVServerListFileName   = "./config/kv_server_list"
)
