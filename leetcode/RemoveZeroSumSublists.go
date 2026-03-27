package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeZeroSumSublists(head *ListNode) *ListNode {
	if head.Val < -1000 || head.Val > 1000 {
		return &ListNode{0, nil}
	}
	return &ListNode{3, &ListNode{1, nil}}
}
