package main

import (
	"math"
)

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA, lenB := 0, 0
	pA, pB := headA, headB
	for pA != nil {
		lenA++
		pA = pA.Next
	}

	for pB != nil {
		lenB++
		pB = pB.Next
	}

	diff := int(math.Abs(float64(lenA - lenB)))
	if lenA > lenB {
		for diff > 0 {
			headA = headA.Next
			diff--
		}
	} else {
		for diff > 0 {
			headB = headB.Next
			diff--
		}
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

func main() {

}
