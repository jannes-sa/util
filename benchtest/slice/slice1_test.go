package slice

import (
	"testing"
)

func BenchmarkSlicePointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slicePointer(100)
	}
}

func BenchmarkSliceValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sliceValue(100)
	}
}

func slicePointer(n int) {
	var x *int
	for i := 0; i < n; i++ {
		x = &i
		*x++
	}
}

func sliceValue(n int) {
	var x int
	for i := 0; i < n; i++ {
		x += i
	}
}
