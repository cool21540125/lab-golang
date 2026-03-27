package leetcode

import (
	"testing"
)

func makeLinkedList1171(lists []int) *ListNode {
	baseNode := &ListNode{0, nil}
	tail := baseNode
	for _, i1 := range lists {
		newNode := ListNode{i1, nil}
		tail.Next = &newNode
		tail = &newNode
	}
	return baseNode.Next
}

// https://leetcode.com/problems/remove-zero-sum-consecutive-nodes-from-linked-list/description/?envType=daily-question&envId=2024-03-13
func TestRemoveZeroSumSublists(t *testing.T) {
	testExamples := []struct {
		head       []int
		targetList []int
	}{
		{
			head:       []int{1, 2, -3, 3, 1},
			targetList: []int{3, 1},
		},
		// {
		// 	head:       []int{1, 2, 3, -3, 4},
		// 	targetList: []int{1, 2, 4},
		// },
		// {
		// 	head:       []int{1, 2, 3, -3, -2},
		// 	targetList: []int{1},
		// },
	}

	for _, example := range testExamples {

		headList := makeLinkedList1171(example.head)
		expectedList := makeLinkedList1171(example.targetList)

		result := removeZeroSumSublists(headList)
		resultList := []int{}
		for {
			resultList = append(resultList, result.Val)
			if result.Next != nil {
				result = result.Next
			} else {
				break
			}
		}

		if len(resultList) != len(example.targetList) {
			t.Errorf("Expected: %v != Result: %v", example.targetList, resultList)
		} else {
			for i := range resultList {
				if resultList[i] != example.targetList[i] {
					t.Errorf("Expected: %v != Result: %v", example.targetList, resultList)
				}
			}
		}

		if result != expectedList {
			t.Errorf("Expected %v, bot got %v", example.targetList, result)
		}
	}
}
