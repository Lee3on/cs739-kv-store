package utils

import (
	"fmt"
	"load_balancer/consts"
)

func GenRaftAddr(id uint64) string {
	return fmt.Sprintf("http://127.0.0.1:%d", consts.RaftPortBase+id-1)
}
