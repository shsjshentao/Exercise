package Algorithm

import (
	"fmt"
	"testing"
)

var listNode11 = &ListNode{2, listNode12}
var listNode12 = &ListNode{5, listNode13}
var listNode13 = &ListNode{3, nil}
var listNode21 = &ListNode{5, listNode22}
var listNode22 = &ListNode{6, listNode23}
var listNode23 = &ListNode{4, nil}

func TestAddTwoNumbers(t *testing.T) {
	newListNode := AddTwoNumbers(listNode21, listNode11)
	for newListNode != nil {
		fmt.Println(newListNode.Val)
		newListNode = newListNode.Next
	}
}
