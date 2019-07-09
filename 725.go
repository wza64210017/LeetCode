package main

import (
	"fmt"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(root *ListNode, k int) []*ListNode {
	ret := make([]*ListNode, 0, k)

	p := root
	length := 0
	for p != nil {
		length++
		p = p.Next
	}

	splitArr := make([]int, 0, k)
	for length > 0 {
		splitNum := int(math.Ceil(float64(length) / float64(k)))
		splitArr = append(splitArr, splitNum)

		k--
		length -= splitNum
	}

	i := 0
	pos := 0
	for root != nil {
		if i == 0 {
			ret = append(ret, root)
		}

		i++
		if i == splitArr[pos] {
			
			i = 0
			pos++
		}

		root = root.Next

	}

	return ret
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	list := &ListNode{
		Val:  arr[0],
		Next: nil,
	}

	tmp := list
	for i := 1; i < len(arr); i++ {
		tmp.Next = &ListNode{
			Val:  arr[i],
			Next: nil,
		}

		tmp = tmp.Next
	}

	ret := splitListToParts(list, 3)
	for _, v := range ret {
		fmt.Println(v.Val)
	}
}
