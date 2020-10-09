package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//func sortList(head *ListNode) *ListNode {
//	var arr []int
//	pq := h(arr)
//	heap.Init(&pq)
//
//	l := head
//	for l != nil {
//		heap.Push(&pq, l.Val)
//		l = l.Next
//	}
//
//	a := head
//	for pq.Len() > 0 {
//		tmp := heap.Pop(&pq).(int)
//		a.Val = tmp
//		a = a.Next
//	}
//
//	return head
//}
//
//type h []int
//
//func (he h) Len() int {
//	return len(he)
//}
//
//func (he h) Swap(i, j int) {
//	he[i], he[j] = he[j], he[i]
//}
//
//func (he h) Less(i, j int) bool {
//	return he[i] < he[j]
//}
//
//func (he *h) Push(i interface{}) {
//	*he = append(*he, i.(int))
//}
//
//func (he *h) Pop() interface{} {
//	length := (*he).Len()
//	if length == 0 {
//		return nil
//	}
//
//	tmp := (*he)[length-1]
//	*he = (*he)[:length-1]
//	return tmp
//}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		fast, slow = fast.Next.Next, slow.Next
	}

	mid := slow.Next
	slow.Next = nil

	left := sortList(head)
	right := sortList(mid)

	ret := new(ListNode)
	h := ret

	for left != nil && right != nil {
		if left.Val < right.Val {
			h.Next, left = left, left.Next
		} else {
			h.Next, right = right, right.Next
		}

		h = h.Next
	}

	if left != nil {
		h.Next = left
	} else {
		h.Next = right
	}

	return ret.Next
}

func main() {
	arr := []int{3, 2, 5, 4, 1}
	list := &ListNode{Val: arr[0], Next: nil}

	tmp := list
	for i := 1; i < len(arr); i++ {
		node := &ListNode{Val: arr[i], Next: nil}
		tmp.Next = node
		tmp = tmp.Next
	}

	l := sortList(list)
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}
