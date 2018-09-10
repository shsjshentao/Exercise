package Algorithm

import "fmt"

func singleNumber(nums []int) []int {
	answer := make([]int, 0, 2)
	store := make(map[int]int)
	for _, v := range nums {
		store[v]++
	}
	fmt.Println(store)
	for k, v := range store {
		if v == 1 {
			answer = append(answer, k)
		}
	}
	return answer
}
