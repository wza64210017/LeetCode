package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	length1 := count(l1)
	length2 := count(l2)

	arr1 := make([]int, 0, length1)
	arr2 := make([]int, 0, length2)

	for l1 != nil {
		arr1 = append(arr1, l1.Val)
		l1 = l1.Next
	}

	for l2 != nil {
		arr2 = append(arr2, l2.Val)
		l2 = l2.Next
	}

	var newL *ListNode

	flow := 0
	sum := 0
	for length1 > 0 && length2 > 0 {
		sum, flow = calSum(arr1[length1-1], arr2[length2-1], flow)
		tmp := &ListNode{
			Val:  sum,
			Next: newL,
		}

		newL = tmp
		length1--
		length2--
	}

	for length1 > 0 {
		sum, flow = calSum(arr1[length1-1], 0, flow)
		tmp := &ListNode{
			Val:  sum,
			Next: newL,
		}

		newL = tmp
		length1--
	}

	for length2 > 0 {
		sum, flow = calSum(0, arr2[length2-1], flow)
		tmp := &ListNode{
			Val:  sum,
			Next: newL,
		}

		newL = tmp
		length2--
	}

	if flow != 0 {
		tmp := &ListNode{
			Val: flow,
			Next: newL,
		}
		newL = tmp
	}

	return newL
}

func count(list *ListNode) int {
	count := 0
	for list != nil {
		count++
		list = list.Next
	}

	return count
}

func calSum(a, b, flow int) (int, int) {
	sum := a + b + flow
	if sum >= 10 {
		flow = sum / 10
		sum %= 10
	} else {
		flow = 0
	}

	return sum, flow
}

func main() {
	arr1 := []int{9, 9, 9, 9}
	list1 := &ListNode{
		Val:  arr1[0],
		Next: nil,
	}

	tmp1 := list1
	for i := 1; i < len(arr1); i++ {
		tmp1.Next = &ListNode{
			Val:  arr1[i],
			Next: nil,
		}

		tmp1 = tmp1.Next
	}

	arr2 := []int{1}
	list2 := &ListNode{
		Val:  arr2[0],
		Next: nil,
	}

	tmp2 := list2
	for i := 1; i < len(arr2); i++ {
		tmp2.Next = &ListNode{
			Val:  arr2[i],
			Next: nil,
		}

		tmp2 = tmp2.Next
	}

	newL := addTwoNumbers(list1, list2)
	for newL != nil {
		fmt.Println(newL.Val)
		newL = newL.Next
	}
}
