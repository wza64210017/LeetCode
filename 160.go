package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	nodeMap := make(map[*ListNode]bool)

	for headA != nil {
		nodeMap[headA] = true
		headA = headA.Next
	}

	for headB != nil {
		if nodeMap[headB] == true {
			return headB
		}

		headB = headB.Next
	}

	return nil
}

func main() {
	arr := []int{2, 6, 4}
	listA := &ListNode{
		Val:  arr[0],
		Next: nil,
	}

	tmp := listA
	for i := 1; i < len(arr); i++ {
		tmp.Next = &ListNode{
			Val:  arr[i],
			Next: nil,
		}

		tmp = tmp.Next
	}

	arr = []int{1, 5}
	listB := &ListNode{
		Val:  arr[0],
		Next: nil,
	}

	tmp = listB
	for i := 1; i < len(arr); i++ {
		tmp.Next = &ListNode{
			Val:  arr[i],
			Next: nil,
		}

		tmp = tmp.Next
	}

	arr = []int{2, 4}
	listC := &ListNode{
		Val:  arr[0],
		Next: nil,
	}

	tmp = listC
	for i := 1; i < len(arr); i++ {
		tmp.Next = &ListNode{
			Val:  arr[i],
			Next: nil,
		}

		tmp = tmp.Next
	}

	//listA.Next = listC
	//listB.Next = listC

	fmt.Println(getIntersectionNode(listA, listB))
}
