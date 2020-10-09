package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	node := reverseList(head.Next) // 最终由第12行返回
	tmp := head.Next
	head.Next = nil
	tmp.Next = head

	// 1 -> 2 -> 3 -> 4 -> null
	// null <- 1 <- 2    3 -> 4 -> null
	// null <- 1 <- 2 <- 3  4 -> null
	// null <- 1 <- 2 <- 3 <- 4
	return node
}

func main() {
	arr := []int{1, 2, 3, 4}
	list := &ListNode{
		Val:  arr[0],
		Next: nil,
	}

	p := list
	for i := 1; i < len(arr); i++ {
		p.Next = &ListNode{
			Val:  arr[i],
			Next: nil,
		}

		p = p.Next
	}

	list = reverseList(list)

	for list != nil {
		//fmt.Println(list.Val)
		list = list.Next
	}
}
