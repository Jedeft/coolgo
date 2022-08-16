package map_slice

import "testing"

// Item 测试结构
type Item struct {
	id  int
	val [1024]byte
}

// BenchmarkIndexStructSlice Cool index获取slice元素值
// BenchmarkIndexStructSlice-16             2583750               462.1 ns/op             0 B/op          0 allocs/op
func BenchmarkIndexStructSlice(b *testing.B) {
	var items [1024]Item
	for i := 0; i < b.N; i++ {
		var tmp int
		for j := 0; j < len(items); j++ {
			tmp = items[j].id
		}
		_ = tmp
	}
}

// BenchmarkRangeIndexStructSlice Cool range通过index获取slice元素值，性能差异不大
// BenchmarkRangeIndexStructSlice-16        2648104               460.2 ns/op             0 B/op          0 allocs/op
func BenchmarkRangeIndexStructSlice(b *testing.B) {
	var items [1024]Item
	for i := 0; i < b.N; i++ {
		var tmp int
		for k := range items {
			tmp = items[k].id
		}
		_ = tmp
	}
}

// BenchmarkRangeStructSlice Bad range 获取slice元素值，性能较差
// 通过range的方式获取value时会将slice中对应下标的元素拷贝到value中，存在一次copy消耗，对value的操作不影响slice数组中原元素属性，且对象越大，性能越差
// BenchmarkRangeStructSlice-16               37783             32630 ns/op               0 B/op          0 allocs/op
func BenchmarkRangeStructSlice(b *testing.B) {
	var items [1024]Item
	for i := 0; i < b.N; i++ {
		var tmp int
		for _, item := range items {
			tmp = item.id
		}
		_ = tmp
	}
}
