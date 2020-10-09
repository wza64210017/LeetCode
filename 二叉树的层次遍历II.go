package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res [][]int

func levelOrderBottom(root *TreeNode) [][]int {
	res = make([][]int, 0)
	dfs(root, 0)
	//for i := 0; i < len(res)/2; i++ {
	//	res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	//}

	fmt.Println(res)
	return res
}

func dfs(root *TreeNode, level int) {
	if root == nil {
		return
	}

	if len(res) == level {
		res = append(res, []int{})
	}

	res[level] = append(res[level], root.Val)

	dfs(root.Left, level+1)
	dfs(root.Right, level+1)
}
