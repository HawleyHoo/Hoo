package h_pkg

import "sync"
//go语言坑之并发访问map
type SafeMap struct {
	sync.RWMutex
	Map map[int]int
}

func NewSafeMap(size int) *SafeMap {
	sm := new(SafeMap)
	sm.Map = make(map[int]int)
	return sm

}

func (sm *SafeMap) ReadMap(key int) int {
	sm.RLock()
	value := sm.Map[key]
	sm.RUnlock()
	return value
}

func (sm *SafeMap) WriteMap(key int, value int) {
	sm.Lock()
	sm.Map[key] = value
	sm.Unlock()
}







