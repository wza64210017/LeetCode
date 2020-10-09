package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	la, lb := calLen(headA), calLen(headB)

	if la > lb {
		headA = move(headA, la-lb)
	} else {
		headB = move(headB, lb-la)
	}

	for headA != nil && headB != nil {
		if headA == headB {
			return headA
		}

		headA = headA.Next
		headB = headB.Next
	}

	return nil
}

func calLen(head *ListNode) int {
	var length int
	p := head
	for p != nil {
		length++
		p = p.Next
	}

	return length
}

func move(head *ListNode, length int) *ListNode {
	for length > 0 {
		head = head.Next
		length--
	}

	return head
}

func main() {

}
