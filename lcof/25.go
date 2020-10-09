package main

//func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
//	p := &ListNode{Val: 0, Next: nil}
//	head := p
//
//	for l1 != nil && l2 != nil {
//		if l1.Val < l2.Val {
//			p.Next = l1
//			l1 = l1.Next
//		} else {
//			p.Next = l2
//			l2 = l2.Next
//		}
//
//		p = p.Next
//	}
//
//	if l1 != nil {
//		p.Next = l1
//	}
//
//	if l2 != nil {
//		p.Next = l2
//	}
//
//	return head.Next
//}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	if l1.Val >= l2.Val {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}

	l1.Next = mergeTwoLists(l1.Next, l2)
	return l1
}

func main() {
	l1 := createList([]int{1, 2, 4})
	l2 := createList([]int{1, 3, 4})
	head := mergeTwoLists(l1, l2)

	printList(head)
}
