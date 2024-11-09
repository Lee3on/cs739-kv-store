package consts

const (
	KVServerListFileName   = "./config/kv_server_list"
	RaftServerListFileName = "./config/raft_server_list"
)

const (
	RaftPortBase         = 5000
	ForceRemoveThreshold = 12
)

const (
	Success       = 0
	InternalError = -1
	KeyNotFound   = 1
	Redirect      = 2
)
