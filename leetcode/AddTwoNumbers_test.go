package leetcode

import (
	"testing"
)

func makeLinkedList(lists []int) *ListNode {
	baseNode := &ListNode{0, nil}
	tail := baseNode
	for _, i1 := range lists {
		newNode := ListNode{i1, nil}
		tail.Next = &newNode
		tail = &newNode
	}
	return baseNode.Next
}

func TestAddTwoNumbers(t *testing.T) {
	testExamples := []struct {
		list1      []int
		list2      []int
		targetList []int
	}{
		{
			list1:      []int{2, 4},
			list2:      []int{5, 6, 4},
			targetList: []int{7, 0, 5},
		},
		{
			list1:      []int{2, 4, 3},
			list2:      []int{5, 6, 4},
			targetList: []int{7, 0, 8},
		},
		{
			list1:      []int{0},
			list2:      []int{0},
			targetList: []int{0},
		},
		{
			list1:      []int{1},
			list2:      []int{9, 0},
			targetList: []int{0, 1},
		},
		{
			list1:      []int{9, 9, 9, 9, 9, 9, 9},
			list2:      []int{9, 9, 9, 9},
			targetList: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	}

	for _, example := range testExamples {

		l1 := makeLinkedList(example.list1)
		l2 := makeLinkedList(example.list2)

		result := addTwoNumbers(l1, l2)

		result_list := []int{}
		for {
			result_list = append(result_list, result.Val)
			if result.Next != nil {
				result = result.Next
			} else {
				break
			}
		}

		if len(result_list) != len(example.targetList) {
			t.Errorf("Expected: %v != Result: %v", example.targetList, result_list)
		} else {
			for i := range result_list {
				if result_list[i] != example.targetList[i] {
					t.Errorf("Expected: %v != Result: %v", example.targetList, result_list)
				}
			}
		}
	}
}
