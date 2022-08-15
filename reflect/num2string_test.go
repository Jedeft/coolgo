package reflect

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

// BenchmarkStrconv Bad
// fmt利用反射达到泛型效果，运行时动态做类型判断，性能损耗相对较大
// BenchmarkStrconv-16     12099537                98.20 ns/op           32 B/op          2 allocs/op
func BenchmarkStrconv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprint(rand.Int())
	}
}

// BenchmarkFmt Cool
// BenchmarkFmt-16         23751020                49.69 ns/op           23 B/op          1 allocs/op
func BenchmarkFmt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strconv.Itoa(rand.Int())
	}
}
