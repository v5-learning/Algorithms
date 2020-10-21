package algorithms

func InsertSort(set ElemSet) {
	length := set.Len()
	for i := 1; i < length; i++ {
		for j := i; j >= 1; j-- {
			if set.Less(j-1, j) {
				set.Swap(j-1, j)
				continue
			}
			break
		}
	}
}
