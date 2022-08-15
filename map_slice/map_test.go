package map_slice

import "testing"

// BenchmarkNoInitSizeMap Bad
// BenchmarkNoInitSizeMap-16         262426              4431 ns/op            5365 B/op         16 allocs/op
func BenchmarkNoInitSizeMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := make(map[int]int)
		for i := 0; i < 100; i++ {
			a[i] = i
		}
		_ = a
	}
}

// BenchmarkInitSizeMap Cool
// BenchmarkInitSizeMap-16           502104              2346 ns/op            2908 B/op          6 allocs/op
// 向make()提供cap会提示在初始化时尝试调整map大小，这将减少将元素添加map时为map重新分配内存，与slice不同的是，map capacity提示并不保证完全的抢占式分配，而是用于估计所需的 hashmap bucket 的数量。因此，在将元素添加到 map 时，甚至在指定 map 容量时，仍可能发生分配。

func BenchmarkInitSizeMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := make(map[int]int, 100)
		for i := 0; i < 100; i++ {
			a[i] = i
		}
		_ = a
	}
}
