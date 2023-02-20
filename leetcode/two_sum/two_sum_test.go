package main

import (
	slices "golang.org/x/exp/slices"
	"testing"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func (n ListNode) GetList() []int {
	res := []int{n.Val}
	for n.Next != nil {
		n = *n.Next
		res = append(res, n.Val)
	}
	return res
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var remember, sum int
	var l_start, l *ListNode
	for l1 != nil || l2 != nil || remember != 0 {
		sum = remember
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		remember = sum / 10
		if l == nil {
			l_start = &ListNode{Val: sum % 10, Next: nil}
			l = l_start
		} else {
			l.Next = &ListNode{Val: sum % 10, Next: nil}
			l = l.Next
		}
	}
	return l_start
}

func TestTwoSum(t *testing.T) {
	var a, b *ListNode
	t.Run("Case 1", func(t *testing.T) {
		a = &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}}
		b = &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}}
		expected_list := []int{7, 0, 8}
		res := addTwoNumbers(a, b).GetList()
		if !slices.Equal(res, expected_list) {
			t.Errorf("Expected %v, got %v\n", expected_list, res)
		}
	})
	t.Run("Case 1", func(t *testing.T) {
		a = &ListNode{Val: 0, Next: nil}
		b = &ListNode{Val: 0, Next: nil}
		expected_list := []int{0}
		res := addTwoNumbers(a, b).GetList()
		if !slices.Equal(res, expected_list) {
			t.Errorf("Expected %v, got %v\n", expected_list, res)
		}
	})
	t.Run("Case 1", func(t *testing.T) {
		a = &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: nil}}}}}}}
		b = &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: nil}}}}
		expected_list := []int{8, 9, 9, 9, 0, 0, 0, 1}
		res := addTwoNumbers(a, b).GetList()
		if !slices.Equal(res, expected_list) {
			t.Errorf("Expected %v, got %v\n", expected_list, res)
		}
	})
}
