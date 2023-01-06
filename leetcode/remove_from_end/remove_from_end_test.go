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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	count_len := 1
	node := head
	for node.Next != nil {
		node = node.Next
		count_len += 1
	}
	start_idx := count_len - n - 1
	node = head
	if start_idx < 0 {
		return head.Next
	} else {
		for i := 0; i < start_idx; i++ {
			node = node.Next
		}
		// Now node is the node to re-link
		node.Next = node.Next.Next
	}
	return head
}

func TestRemoveFromEnd(t *testing.T) {
	v := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{4, &ListNode{Val: 5, Next: nil}}}}}
	q := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}}
	t.Run("Print Link List", func(t *testing.T) {
		expected_list := []int{1, 2, 3, 4, 5}
		res := v.GetList()
		if !slices.Equal(res, expected_list) {
			t.Errorf("Expected %v, got %v\n", expected_list, res)
		}
	})

	t.Run("RemoveNthFromEnd", func(t *testing.T) {
		remove_idx := 2
		expected_list := []int{2}
		q := removeNthFromEnd(q, remove_idx)
		res := q.GetList()
		if !slices.Equal(res, expected_list) {
			t.Errorf("Expected %v, got %v\n", expected_list, res)
		}
	})
}
