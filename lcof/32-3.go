package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	dfs(root, 0, &ret)
	return ret
}

func dfs(root *TreeNode, level int, ret *[][]int) {
	if root == nil {
		return
	}

	if level == len(*ret) {
		*ret = append(*ret, make([]int, 0))
	}

	if level%2 == 1 {
		(*ret)[level] = append([]int{root.Val}, (*ret)[level]...)
	} else {
		(*ret)[level] = append((*ret)[level], root.Val)
	}

	dfs(root.Left, level+1, ret)
	dfs(root.Right, level+1, ret)
}
