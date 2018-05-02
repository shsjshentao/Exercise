package Algorithm

type ListNode struct {
	Val  int
	Next *ListNode
}


func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := new(ListNode)
	p := l1
	q := l2
	m := l3
	carry := 0
	for {
		if p == nil && q == nil {
			break
		}
		var tmp1, tmp2 int
		if p != nil {
			tmp1 = p.Val
		} else {
			tmp1 = 0
		}
		if q != nil {
			tmp2 = q.Val
		} else {
			tmp2 = 0
		}
		sum := tmp1 + tmp2 + carry
		carry = 0
		carry = sum / 10
		sum = sum % 10
		m.Next = &ListNode{Val: sum}
		m = m.Next
		if p != nil {
			p = p.Next
		}
		if q != nil {
			q = q.Next
		}
	}
	if carry > 0 {
		m.Next = &ListNode{Val: carry}
	}
	return l3.Next
}
