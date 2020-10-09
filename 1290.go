package main

type node struct {
	Val  int
	Next *node
}

func getDecimalValue(head *node) int {
	num := 0
	for head != nil {
		num = num << 1
		num += head.Val
		head = head.Next
	}

	return num
}


