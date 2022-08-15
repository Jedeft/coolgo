// golang小知识
// 栈内存：函数中申请一个对象，函数执行结束，则自动回收。
// 分配回收内存开销仅2个CPU指令：PUSH/POP，在栈上分配内存消耗仅是数据copy至内存的时间，所以分配内存效率非常高
// 堆内存：函数中申请一个对象，函数执行结束后在某一个时间点回收
// 堆分配额外开销主要在于垃圾回收，golang采用三色标记法进行垃圾回收，清理过程主分为标记、清除两个阶段。典型耗时在标记期间需要暂停程序（Stop the world，STW，耗时1ms左右），标记结束之后，用户程序才可以继续执行。

package reflect

import (
	"reflect"
	"testing"
)

// BenchmarkDeleteSliceElms Bad
// BenchmarkDeleteSliceElms-16              1463187               823.7 ns/op           296 B/op         16 allocs/op
func BenchmarkDeleteSliceElms(b *testing.B) {
	slice := []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	elms := []interface{}{uint64(1), uint64(3), uint64(5), uint64(7), uint64(9)}
	for i := 0; i < b.N; i++ {
		_ = DeleteSliceElms(slice, elms...)
	}
}

// BenchmarkDeleteU64liceElms Cool
// BenchmarkDeleteU64liceElms-16            9343531               124.5 ns/op            80 B/op          1 allocs/op
func BenchmarkDeleteU64liceElms(b *testing.B) {
	slice := []uint64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	elms := []uint64{1, 3, 5, 7, 9}
	for i := 0; i < b.N; i++ {
		_ = DeleteU64liceElms(slice, elms...)
	}
}

// DeleteSliceElms 从切片中过滤指定元素。注意：不修改原切片。
// 虽然泛化了使用，但不太存在对所有类型切片进行过滤的场景，通过反射带来的便利性并没有得到很好的体现，相反带来了高昂的性能负担
// map的key是interface{}，编译期间很难确定具体类型，将发生变量逃逸，初始化堆内存分配，性能开销大
func DeleteSliceElms(i interface{}, elms ...interface{}) interface{} {
	// 构建 map set。
	m := make(map[interface{}]struct{}, len(elms))
	for _, v := range elms {
		m[v] = struct{}{}
	}
	// 创建新切片，过滤掉指定元素。
	v := reflect.ValueOf(i)
	t := reflect.MakeSlice(reflect.TypeOf(i), 0, v.Len())
	for i := 0; i < v.Len(); i++ {
		if _, ok := m[v.Index(i).Interface()]; !ok {
			t = reflect.Append(t, v.Index(i))
		}
	}
	return t.Interface()
}

// DeleteU64liceElms 从 []uint64 过滤指定元素。注意：不修改原切片。
// 出入参指定类型，如果存在其他类型过滤的诉求，那么建议copy此函数重新实现
func DeleteU64liceElms(i []uint64, elms ...uint64) []uint64 {
	// 构建 map set。
	m := make(map[uint64]struct{}, len(elms))
	for _, v := range elms {
		m[v] = struct{}{}
	}
	// 创建新切片，过滤掉指定元素。
	t := make([]uint64, 0, len(i))
	for _, v := range i {
		if _, ok := m[v]; !ok {
			t = append(t, v)
		}
	}
	return t
}
