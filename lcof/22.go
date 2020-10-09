package main

func getKthFromEnd(head *ListNode, k int) *ListNode {
	if k == 0 {
		return head
	}

	prA := head
	prB := head
	for k > 0 && prB != nil {
		prB = prB.Next
		k--
	}

	for prB != nil {
		prA = prA.Next
		prB = prB.Next
	}

	return prA
}

func main() {
	head := createList([]int{1, 2, 3, 4, 5})
	head = getKthFromEnd(head, 0)
	printList(head)
}
