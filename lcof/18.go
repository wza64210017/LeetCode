package main

func deleteNode(head *ListNode, val int) *ListNode {
	if head.Val == val {
		return head.Next
	}

	p := head
	for p.Next.Val != val {
		p = p.Next
	}

	p.Next = p.Next.Next

	return head
}

func main() {
	list := createList([]int{4, 5, 1, 9})
	list = deleteNode(list, 9)
	printList(list)
}
