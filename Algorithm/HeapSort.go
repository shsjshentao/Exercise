package Algorithm

import (
	"fmt"
)

func heapAdjust(heapArray []int, pos int, length int) { //调整大顶堆
	temp := heapArray[pos]
	for k := 2*pos + 1; k < length; k = 2*k + 1 {
		if k+1 < length && heapArray[k] < heapArray[k+1] {
			k++
		}
		if heapArray[pos] < heapArray[k] {
			heapArray[pos] = heapArray[k]
			pos = k
			heapArray[pos] = temp
		} else {
			break
		}

	}

}

func HeapInit(heapArray []int) { //构造初始堆
	for i := len(heapArray)/2 - 1; i >= 0; i-- { //从第一个非叶子结点从下至上，从右至左调整结构
		heapAdjust(heapArray, i, len(heapArray))
	}
	fmt.Println(heapArray)
}

func HeapSort(heapArray []int) {
	HeapInit(heapArray)
	for j := len(heapArray) - 1; j >= 0; j-- {
		heapArray[0], heapArray[j] = heapArray[j], heapArray[0]
		heapAdjust(heapArray, 0, j)
	}
	fmt.Println(heapArray)
}
