package utils

import (
	"cs739-kv-store/consts"
	"fmt"
)

func GenRaftAddr(id uint64) string {
	return fmt.Sprintf("http://127.0.0.1:%d", consts.RaftPortBase+id-1)
}
