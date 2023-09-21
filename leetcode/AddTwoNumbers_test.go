package leetcode

import (
	"testing"
)

func TestAddTwoNumbers1(t *testing.T) {
	list1 := []int{2, 4}
	list2 := []int{5, 6, 4}
	target := []int{7, 0, 5}

	base1 := &ListNode{0, nil}
	tail := base1
	for _, i1 := range list1 {
		newNode := ListNode{i1, nil}
		tail.Next = &newNode
		tail = &newNode
	}
	l1 := base1.Next

	base2 := &ListNode{0, nil}
	tail = base2
	for _, i2 := range list2 {
		newNode := ListNode{i2, nil}
		tail.Next = &newNode
		tail = &newNode
	}
	l2 := base2.Next

	baseTarget := &ListNode{0, nil}
	tail = baseTarget
	for _, i3 := range target {
		newNode := ListNode{i3, nil}
		tail.Next = &newNode
		tail = &newNode
	}
	expected := baseTarget.Next

	result := addTwoNumbers(l1, l2)
	if result.Val != expected.Val {
		t.Errorf("result.Val: %v != expected.Val: %v", result.Val, expected.Val)
	}
}

func TestAddTwoNumbers2(t *testing.T) {
	list1 := []int{2, 4, 3}
	list2 := []int{5, 6, 4}
	target := []int{7, 0, 8}

	base1 := &ListNode{0, nil}
	tail := base1
	for _, i1 := range list1 {
		newNode := ListNode{i1, nil}
		tail.Next = &newNode
		tail = &newNode
	}
	l1 := base1.Next

	base2 := &ListNode{0, nil}
	tail = base2
	for _, i2 := range list2 {
		newNode := ListNode{i2, nil}
		tail.Next = &newNode
		tail = &newNode
	}
	l2 := base2.Next

	baseTarget := &ListNode{0, nil}
	tail = baseTarget
	for _, i3 := range target {
		newNode := ListNode{i3, nil}
		tail.Next = &newNode
		tail = &newNode
	}
	expected := baseTarget.Next

	result := addTwoNumbers(l1, l2)
	if result.Val != expected.Val {
		t.Errorf("result.Val: %v != expected.Val: %v", result.Val, expected.Val)
	}
}

func TestAddTwoNumbers3(t *testing.T) {
	list1 := []int{0}
	list2 := []int{0}
	target := []int{0}

	base1 := &ListNode{0, nil}
	tail := base1
	for _, i1 := range list1 {
		newNode := ListNode{i1, nil}
		tail.Next = &newNode
		tail = &newNode
	}
	l1 := base1.Next

	base2 := &ListNode{0, nil}
	tail = base2
	for _, i2 := range list2 {
		newNode := ListNode{i2, nil}
		tail.Next = &newNode
		tail = &newNode
	}
	l2 := base2.Next

	baseTarget := &ListNode{0, nil}
	tail = baseTarget
	for _, i3 := range target {
		newNode := ListNode{i3, nil}
		tail.Next = &newNode
		tail = &newNode
	}
	expected := baseTarget.Next

	result := addTwoNumbers(l1, l2)
	if result.Val != expected.Val {
		t.Errorf("result.Val: %v != expected.Val: %v", result.Val, expected.Val)
	}
}

func TestAddTwoNumbers4(t *testing.T) {
	list1 := []int{9, 9, 9, 9, 9, 9, 9}
	list2 := []int{9, 9, 9, 9}
	target := []int{8, 9, 9, 9, 0, 0, 0, 1}

	base1 := &ListNode{0, nil}
	tail := base1
	for _, i1 := range list1 {
		newNode := ListNode{i1, nil}
		tail.Next = &newNode
		tail = &newNode
	}
	l1 := base1.Next

	base2 := &ListNode{0, nil}
	tail = base2
	for _, i2 := range list2 {
		newNode := ListNode{i2, nil}
		tail.Next = &newNode
		tail = &newNode
	}
	l2 := base2.Next

	baseTarget := &ListNode{0, nil}
	tail = baseTarget
	for _, i3 := range target {
		newNode := ListNode{i3, nil}
		tail.Next = &newNode
		tail = &newNode
	}
	expected := baseTarget.Next

	result := addTwoNumbers(l1, l2)
	if result.Val != expected.Val {
		t.Errorf("result.Val: %v != expected.Val: %v", result.Val, expected.Val)
	}
}
