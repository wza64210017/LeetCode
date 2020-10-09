package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	dfs(root, &ret, 0)

	return ret
}

func dfs(root *TreeNode, ret *[][]int, level int) {
	if root == nil {
		return
	}

	if len(*ret) == level {
		*ret = append(*ret, make([]int, 0))
	}

	(*ret)[level] = append((*ret)[level], root.Val)
	dfs(root.Left, ret, level+1)
	dfs(root.Right, ret, level+1)
}
