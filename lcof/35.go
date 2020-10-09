package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	m := make(map[*Node]*Node)

	newH := &Node{Val: head.Val}
	m[head] = newH

	for p1, p2 := head, newH; p1 != nil; p1, p2 = p1.Next, p2.Next {
		node := &Node{Val: p1.Val}
		m[p1] = node
		p2.Next = node
	}

	for p1, p2 := head, newH; p1 != nil; p1, p2 = p1.Next, p2.Next {
		if p1.Random != nil {
			p2.Random = m[p1.Random]
		}
	}

	return newH
}
