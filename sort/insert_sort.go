package algorithms

func InsertSort(tmp ElemSet) {
	length := tmp.Len()
	for i := 1; i < length; i++ {
		for j := i; j >= 1; j-- {
			if tmp.Less(j-1, j) {
				tmp.Swap(j-1, j)
				continue
			}
			break
		}
	}
}
