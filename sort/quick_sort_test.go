package algorithms

import (
	"fmt"
	"sort"
	"testing"
)

var n100 ElemSet
var n1000 ElemSet
var n10000 ElemSet
var n100000 ElemSet

func init() {
	n100 = benchmarkIntn(100)
	n1000 = benchmarkIntn(1000)
	n10000 = benchmarkIntn(10000)
	n100000 = benchmarkIntn(100000)
}

func TestIntQuickSort(t *testing.T) {
	//tmp := []int{9, 8, 7, 6, 4, 1, 4}
	tmp := []int{1, 3, 1, 4, 1, 5}
	set := intSet(tmp)
	fmt.Println(tmp)
	QuickSort(set)
	fmt.Println(tmp)

	tmp = []int{9, 8, 7, 6, 4, 1, 4, 1, 1, 1, 1, 1, 43, 7, 8, 3, 2, 32, 32, 31431412, 42, 23423423}
	set = intSet(tmp)
	fmt.Println(tmp)
	QuickSort(set)
	fmt.Println(tmp)

	tmp = []int{9, 8, 7, 6, 4, 1, 4, 1, 1, 1, 1, 1, 43, 7, 8, 3, 2, 32, 32, 31431412, 42, 23423423, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 3, 3, 3, 3, 4444444444444, 1, 1, 1, 12, 2, 3, 3, 4, 3, 23, 4134141, 143124124, 324, 1234123, 4124, 124, 12, 42, 412, 412, 412, 4, 124, 124, 124, 1, 42, 34, 2314, 1234, 2314, 34, 35, 465, 56, 5, 65, 76, 76, 6, 6, 567, 567, 56}
	set = intSet(tmp)
	fmt.Println(tmp)
	QuickSort(set)
	fmt.Println(tmp)

}

func BenchmarkIntQuick100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSort(n100)
		//sort.Sort(set)
	}
}

func BenchmarkIntQuick1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSort(n1000)
		//sort.Sort(set)
	}
}

func BenchmarkIntQuick10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort.Sort(set)
		QuickSort(n10000)
	}
}

func BenchmarkIntQuick100000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//sort.Sort(set)
		QuickSort(n100000)
	}
}
