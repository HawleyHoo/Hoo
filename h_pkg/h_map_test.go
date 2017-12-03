package h_pkg

import "testing"

func TestSafeMap_ReadMap(t *testing.T) {
	safeMap := NewSafeMap(10)

	for i := 0; i < 200000; i++ {
		go safeMap.WriteMap(i, i)
		go safeMap.ReadMap(i)
	}
}
