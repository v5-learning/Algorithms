package algorithms

func QuickSort(set ElemSet) {
	quickSort(set, 0, set.Len()-1)
}

func quickSort(set ElemSet, start, end int) {
	if start >= end {
		return
	}

	//sep := partition(start, end)
	pos := swap(set, start, end)
	// [start,pos]
	quickSort(set, start, pos-1)

	// [pos+1,end]
	quickSort(set, pos+1, end)
}

func partition(start, end int) int {
	return start
}

func swap(set ElemSet, left, right int) int {
	sep := left

	for left < right {
		for left < right && set.Less(right, sep) { // right >= sep
			right--
		}

		set.Swap(left, right)
		sep = right

		for left < right && set.Less(sep, left) { // sep >= left
			left++
		}

		set.Swap(left, right)
		//fmt.Println(reflect.ValueOf(set).Index(left), "交换", reflect.ValueOf(set).Index(right), "piovt:", reflect.ValueOf(set).Index(sep))
		sep = left
	}

	//return sep
	/*
		if set.Less(left, sep) {
			left--
		}

		set.Swap(left, sep)
		sep = left
	*/

	//pivot := reflect.ValueOf(set).Index(sep)
	//fmt.Println("pivot:", sep, "--->", pivot)
	//fmt.Println("before:", set)
	/*
		for left < right {
			for left < right && set.Less(right, sep) {
				//fmt.Println(reflect.ValueOf(set).Index(right), ">=", reflect.ValueOf(set).Index(sep), "right右移", right-1, "<---", right)
				right--
			}

			if left < right {
				//fmt.Println(reflect.ValueOf(set).Index(right), "<", reflect.ValueOf(set).Index(sep), "交换", sep)
				set.Swap(right, sep)
				//fmt.Println(reflect.ValueOf(set).Index(right), ">", reflect.ValueOf(set).Index(sep), "交换", right)
				sep = right

			}

			for left < right && set.Less(sep, left) {
				//fmt.Println(reflect.ValueOf(set).Index(sep), ">=", reflect.ValueOf(set).Index(left), "left左移", left, "-->", left+1)
				left++

			}

			if left < right {
				//fmt.Println(reflect.ValueOf(set).Index(left), ">", reflect.ValueOf(set).Index(sep), "交换", sep)
				set.Swap(left, sep)
				//fmt.Println(reflect.ValueOf(set).Index(sep), ">", reflect.ValueOf(set).Index(left), "交换", left)
				sep = left
			}

				if set.Less(right, sep) { // sep < end
					fmt.Println(reflect.ValueOf(set).Index(right), ">=", reflect.ValueOf(set).Index(sep), "right右移", right-1, "<---", right)
					right--
					continue
				} else {
					fmt.Println(reflect.ValueOf(set).Index(right), "<", reflect.ValueOf(set).Index(sep), "交换", sep)
					set.Swap(right, sep)

					fmt.Println(reflect.ValueOf(set).Index(right), ">", reflect.ValueOf(set).Index(sep), "交换", right)
					sep = right
				}

				if set.Less(sep, left) { // sep > start
					fmt.Println(reflect.ValueOf(set).Index(sep), ">=", reflect.ValueOf(set).Index(left), "left左移", left, "-->", left+1)
					left++
					continue
				} else {
					fmt.Println(reflect.ValueOf(set).Index(left), ">", reflect.ValueOf(set).Index(sep), "交换", sep)
					set.Swap(left, sep)
					fmt.Println(reflect.ValueOf(set).Index(sep), ">", reflect.ValueOf(set).Index(left), "交换", left)
					sep = left
				}
		}

		//reflect.ValueOf(set).Index(left).Set(pivot)
		//fmt.Println("swap after", set)
		//fmt.Println("sep:", sep)
		return sep
	*/
	return sep
}
