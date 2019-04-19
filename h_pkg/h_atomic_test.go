package main

import (
	"testing"
)

const m1 = 0x5555555555555555
const m2 = 0x3333333333333333
const m4 = 0x0f0f0f0f0f0f0f0f
const h01 = 0x0101010101010101

func Popcnt(x uint64) uint64 {
	x -= (x >> 1) & m1
	x = (x & m2) + ((x >> 2) & m2)
	x = (x + (x >> 4)) & m4
	return (x * h01) >> 56
}

func BenchmarkPopcnt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := i
		x -= (x >> 1) & m1
		x = (x & m2) + ((x >> 2) & m2)
		x = (x + (x >> 4)) & m4
		_ = (x * h01) >> 56
	}
}

//type manager struct {
//	sync.Mutex
//	agents int
//}
//
//func BenchmarkManagerLock(b *testing.B)  {
//	m := new(manager)
//	b.ReportAllocs()
//	b.RunParallel(func(pb *testing.PB) {
//		for pb.Next()  {
//			m.Lock()
//			m.agents = 100
//			m.Unlock()
//		}
//	})
//}

//func BenchmarkNewNoaway(b *testing.B) {
//
//}
