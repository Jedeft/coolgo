package map_slice

import (
	"testing"
)

// BenchmarkNoInitCapSlice Bad
// BenchmarkNoInitCapSlice-16       3232816               383.4 ns/op          2040 B/op          8 allocs/op
// 初始化slice时尽量指定size/cap大小，避免元素数量超过cap值时通过复制来调整slice大小
func BenchmarkNoInitCapSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := make([]int, 0)
		for i := 0; i < 100; i++ {
			a = append(a, i)
		}
		_ = a
	}
}

// BenchmarkInitCapSlice Cool
// BenchmarkInitCapSlice-16        22633447                53.05 ns/op            0 B/op          0 allocs/op
func BenchmarkInitCapSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		a := make([]int, 0, 100)
		for i := 0; i < 100; i++ {
			a = append(a, i)
		}
		_ = a
	}
}
