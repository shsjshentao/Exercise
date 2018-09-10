package Algorithm

func QuickSort(quickArray []int) []int {
	return compare(quickArray, 0, len(quickArray)-1)
}

func compare(quickArray []int, left, right int) []int {
	if left > right {
		return quickArray
	}
	base := left
	i := left
	j := right
	for i < j {
		for i < j && quickArray[j] > quickArray[base] {
			j--
		}
		for i < j && quickArray[i] < quickArray[base] {
			i++
		}
		if i < j {
			quickArray[i], quickArray[j] = quickArray[j], quickArray[i]
		}
	}
	quickArray[base], quickArray[i] = quickArray[i], quickArray[base]
	compare(quickArray, left, i-1)
	compare(quickArray, i+1, right)
	return quickArray
}
