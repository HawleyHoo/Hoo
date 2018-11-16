package h_pkg

import (
	"testing"

	"fmt"
)

func TestMutexTs(t *testing.T) {
	//MutexTs()
	var aa float32 = 0
	if aa == 0.0 {
		fmt.Println("0 == 0.0")
	} else {
		fmt.Println("0 != 0.0")
	}
	//SyncWaitGroup()
}
