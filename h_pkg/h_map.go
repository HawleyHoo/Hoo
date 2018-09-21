package h_pkg

import "sync"

/*
RWMutex（读写锁）
RWMutex 是单写多读锁，该锁可以加多个读锁或者一个写锁
读锁占用的情况下会阻止写，不会阻止读，多个 goroutine 可以同时获取读锁
写锁会阻止其他 goroutine（无论读和写）进来，整个锁由该 goroutine 独占
适用于读多写少的场景
Lock() 和 Unlock()
Lock() 加写锁，Unlock() 解写锁
如果在加写锁之前已经有其他的读锁和写锁，则 Lock() 会阻塞直到该锁可用，为确保该锁可用，已经阻塞的 Lock() 调用会从获得的锁中排除新的读取器，即写锁权限高于读锁，有写锁时优先进行写锁定
在 Lock() 之前使用 Unlock() 会导致 panic 异常
RLock() 和 RUnlock()
RLock() 加读锁，RUnlock() 解读锁
RLock() 加读锁时，如果存在写锁，则无法加读锁；当只有读锁或者没有锁时，可以加读锁，读锁可以加载多个
RUnlock() 解读锁，RUnlock() 撤销单词 RLock() 调用，对于其他同时存在的读锁则没有效果
在没有读锁的情况下调用 RUnlock() 会导致 panic 错误
RUnlock() 的个数不得多余 RLock()，否则会导致 panic 错误
*/
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
