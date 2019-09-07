package algorithms

type ElemSet interface {
	Len() int
	Less(int, int) bool
	Swap(int, int)
}

type intSet []int

func (set intSet) Len() int {
	return len(set)
}

func (set intSet) Less(i, j int) bool {
	return set[i] >= set[j]
}

func (set intSet) Swap(i, j int) {
	set[i], set[j] = set[j], set[i]
}

type floatSet []float64

func (set floatSet) Len() int {
	return len(set)
}

func (set floatSet) Less(i, j int) bool {
	return set[i] >= set[j]
}

func (set floatSet) Swap(i, j int) {
	set[i], set[j] = set[j], set[i]
}
