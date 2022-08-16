package map_slice

import (
	"math/rand"
	"testing"
	"time"
)

// genRandomIntSlice 生成指定长度的随机 []int 切片
func genRandomIntSlice(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

// BenchmarkIndexIntSlice 对于[]int，不论是index下标遍历取值还是range，遍历性能没有区别
// BenchmarkIndexIntSlice-16        5214588               228.1 ns/op             0 B/op          0 allocs/op
func BenchmarkIndexIntSlice(b *testing.B) {
	nums := genRandomIntSlice(1024)
	for i := 0; i < b.N; i++ {
		var tmp int
		for k := 0; k < len(nums); k++ {
			tmp = nums[k]
		}
		_ = tmp
	}
}

// BenchmarkRangeIntSlice 对于[]int，不论是index下标遍历取值还是range，遍历性能没有区别
// BenchmarkRangeIntSlice-16        5228493               226.8 ns/op             0 B/op          0 allocs/op
func BenchmarkRangeIntSlice(b *testing.B) {
	nums := genRandomIntSlice(1024)
	for i := 0; i < b.N; i++ {
		var tmp int
		for _, num := range nums {
			tmp = num
		}
		_ = tmp
	}
}
