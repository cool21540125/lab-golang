package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	baseNode := &ListNode{0, nil}
	tail := baseNode

	ten := 0
	rem := 0
	n1, n2 := 0, 0
	for {
		if l1 == nil && l2 == nil {
			break
		}

		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		} else {
			n1 = 0
		}

		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		} else {
			n2 = 0
		}

		summ := n1 + n2 + ten
		if summ >= 10 {
			ten = 1
			rem = summ - 10
		} else {
			ten = 0
			rem = summ
		}

		newNode := ListNode{rem, nil}
		tail.Next = &newNode
		tail = &newNode
	}
	if ten > 0 {
		newNode := ListNode{1, nil}
		tail.Next = &newNode
		tail = &newNode
	}
	return baseNode.Next
}
