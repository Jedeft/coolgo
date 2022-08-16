package map_slice

import "testing"

// genItems 生成指定长度 []*Item 切片
func genItems(n int) []*Item {
	items := make([]*Item, 0, n)
	for i := 0; i < n; i++ {
		items = append(items, &Item{id: i})
	}
	return items
}

// BenchmarkIndexPointer 使用指针以后，两种方式获取元素的性能差异不大
// BenchmarkIndexPointer-16                  814461              1472 ns/op               0 B/op          0 allocs/op
func BenchmarkIndexPointer(b *testing.B) {
	items := genItems(1024)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var tmp int
		for k := 0; k < len(items); k++ {
			tmp = items[k].id
		}
		_ = tmp
	}
}

// BenchmarkIndexPointer 使用指针以后，两种方式获取元素的性能差异不大，并且value值相当于是指向元素的指针拷贝，指向的是同一个对象，支持对slice中的元素进行修改
// BenchmarkRangePointer-16                  794662              1353 ns/op               0 B/op          0 allocs/op
// 虽然使用指针可以便捷的对元素进行修改，但在slice中更推荐使用struct，slice元素存储的是指针，指针又单独指向一个内存区域，使用结构体可以直接省去这部分开销
func BenchmarkRangePointer(b *testing.B) {
	items := genItems(1024)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var tmp int
		for _, item := range items {
			tmp = item.id
		}
		_ = tmp
	}
}
