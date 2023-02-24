package linklist

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

func CreateSinglyLListFromArr(arr []int) *ListNode {
	var l_start, l *ListNode
	for i := 0; i < len(arr); i++ {
		if l_start == nil {
			l_start = &ListNode{Val: arr[i], Next: nil}
			l = l_start
		} else {
			l.Next = &ListNode{Val: arr[i], Next: nil}
			l = l.Next
		}
	}
	return l_start
}
