package string

import (
	"fmt"
	"testing"
)

// BenchmarkJoinStrWithSprintf Bad
// BenchmarkJoinStrWithSprintf-16           8426418               141.5 ns/op            64 B/op          4 allocs/op
// 底层实现使用了反射，性能上有所损耗，当拼接字符串涉及类型转换且数量较少（<=5）推荐使用此方法
func BenchmarkJoinStrWithSprintf(b *testing.B) {
	s1, s2, s3 := "foo", "bar", "baz"
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s%s%s", s1, s2, s3)
	}
}

// BenchmarkJoinStrWithOperator Cool
// BenchmarkJoinStrWithOperator-16         70130563                17.03 ns/op            0 B/op          0 allocs/op
// 若拼接字符串不涉及类型转换，且数量较少（<=5），出于性能考虑推荐使用此方法
func BenchmarkJoinStrWithOperator(b *testing.B) {
	s1, s2, s3 := "foo", "bar", "baz"
	for i := 0; i < b.N; i++ {
		_ = s1 + s2 + s3
	}
}
