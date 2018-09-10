package Algorithm

import "testing"

func TestHeapSort(t *testing.T) {
	a := []int{0, 10, 19, 24, 61, 5, 121, 9, 11, 34, 21, 22}
	//a := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	HeapSort(a)
}
