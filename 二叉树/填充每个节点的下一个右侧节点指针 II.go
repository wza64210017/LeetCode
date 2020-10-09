package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

var nextArr [][]*Node

func connect(root *Node) *Node {
	nextArr = make([][]*Node, 0)
	dfs(root, 0)

	for _, arr := range nextArr {
		for i := len(arr) - 1; i > 0; i-- {
			arr[i-1].Next = arr[i]
		}

		arr[len(arr)-1].Next = nil
	}

	return root
}

func dfs(root *Node, level int) {
	if root == nil {
		return
	}

	if level == len(nextArr) {
		nextArr[level] = make([]*Node, 0)
	}

	nextArr[level] = append(nextArr[level], root)
	dfs(root.Left, level+1)
	dfs(root.Right, level+1)
}
