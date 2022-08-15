package string

import (
	"strings"
	"testing"
)

func BenchmarkConvert(b *testing.B) {
	var buf strings.Builder
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Write([]byte("Hello world"))
	}
}

func BenchmarkNoConvert(b *testing.B) {
	var buf strings.Builder
	b.ResetTimer()
	data := []byte("Hello world")
	for i := 0; i < b.N; i++ {
		buf.Write(data)
	}
}
