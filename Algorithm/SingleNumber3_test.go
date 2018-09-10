package Algorithm

import (
	"fmt"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	a := []int{1, 2, 1, 3, 2, 5}
	fmt.Println(singleNumber(a))
}
