package string

import (
	"bytes"
	"strings"
	"testing"
)

// BenchmarkJoinStrWithStringsJoin 性能相近
// BenchmarkJoinStrWithStringsJoin-16              35743003                33.13 ns/op           16 B/op          1 allocs/op
func BenchmarkJoinStrWithStringsJoin(b *testing.B) {
	s1, s2, s3 := "foo", "bar", "baz"
	for i := 0; i < b.N; i++ {
		_ = strings.Join([]string{s1, s2, s3}, "")
	}
}

// BenchmarkJoinStrWithStringsBuilder 性能相近，推荐
// BenchmarkJoinStrWithStringsBuilder-16           31323195                36.66 ns/op           24 B/op          2 allocs/op
func BenchmarkJoinStrWithStringsBuilder(b *testing.B) {
	s1, s2, s3 := "foo", "bar", "baz"
	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		_, _ = builder.WriteString(s1)
		_, _ = builder.WriteString(s2)
		_, _ = builder.WriteString(s3)
	}
}

// BenchmarkJoinStrWithStringsBuilderPreAlloc Cool，支持提前分配内存大小，性能优化明显
// BenchmarkJoinStrWithStringsBuilderPreAlloc-16           61136556                18.66 ns/op           16 B/op          1 allocs/op
func BenchmarkJoinStrWithStringsBuilderPreAlloc(b *testing.B) {
	s1, s2, s3 := "foo", "bar", "baz"
	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		builder.Grow(9)
		_, _ = builder.WriteString(s1)
		_, _ = builder.WriteString(s2)
		_, _ = builder.WriteString(s3)
	}
}

// BenchmarkJoinStrWithBytesBuffer 性能相近
// BenchmarkJoinStrWithBytesBuffer-16              37176975                31.05 ns/op           64 B/op          1 allocs/op
func BenchmarkJoinStrWithBytesBuffer(b *testing.B) {
	s1, s2, s3 := "foo", "bar", "baz"
	for i := 0; i < b.N; i++ {
		var buffer bytes.Buffer
		_, _ = buffer.WriteString(s1)
		_, _ = buffer.WriteString(s2)
		_, _ = buffer.WriteString(s3)
	}
}

// BenchmarkJoinStrWithByteSlice 性能相近
// BenchmarkJoinStrWithByteSlice-16                34461618                33.33 ns/op           24 B/op          2 allocs/op
func BenchmarkJoinStrWithByteSlice(b *testing.B) {
	s1, s2, s3 := "foo", "bar", "baz"
	for i := 0; i < b.N; i++ {
		var bys []byte
		bys = append(bys, s1...)
		bys = append(bys, s2...)
		_ = append(bys, s3...)
	}
}

// BenchmarkJoinStrWithByteSlicePreAlloc 性能最好，但需要提前预知字符串大小，易用性较差，不推荐使用
// BenchmarkJoinStrWithByteSlicePreAlloc-16        685866504                1.742 ns/op           0 B/op          0 allocs/op
func BenchmarkJoinStrWithByteSlicePreAlloc(b *testing.B) {
	s1, s2, s3 := "foo", "bar", "baz"
	for i := 0; i < b.N; i++ {
		bys := make([]byte, 0, 9)
		bys = append(bys, s1...)
		bys = append(bys, s2...)
		_ = append(bys, s3...)
	}
}
