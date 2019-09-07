package algorithms

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestIntInsert(t *testing.T) {
	tmp := []int{4, 6, 3, 1, 7, 3, 3, 3, 3, 1, 3, 34, 4, 56, 745, 0}
	set := intSet(tmp)
	fmt.Println(tmp)
	InsertSort(set)
	fmt.Println(tmp)
}

func TestFloatInsert(t *testing.T) {
	tmp := []float64{4.0, 6.1, 6.2, 1.1, 7.0}
	set := floatSet(tmp)
	fmt.Println(tmp)
	InsertSort(set)
	fmt.Println(tmp)
}

func BenchmarkIntInsert100(b *testing.B) {
	set := benchmarkIntn(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		InsertSort(set)
	}
	set = nil
}

func BenchmarkIntInsert1000(b *testing.B) {
	set := benchmarkIntn(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		InsertSort(set)
	}
	set = nil
}

func BenchmarkIntInsert10000(b *testing.B) {
	set := benchmarkIntn(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		InsertSort(set)
	}
	set = nil
}

func BenchmarkIntInsert100000(b *testing.B) {
	set := benchmarkIntn(100000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		InsertSort(set)
	}
	set = nil
}

func benchmarkIntn(n int) ElemSet {
	tmp := make([]int, n)
	for i := 0; i < n; i++ {
		tmp[i] = rand.Int()
	}
	return intSet(tmp)
}
