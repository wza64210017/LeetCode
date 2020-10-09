package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	dfs(root, &ret)

	return ret
}

func dfs(root *TreeNode, ret *[]int) {
	if root == nil {
		return
	}

	dfs(root.Left, ret)
	*ret = append(*ret, root.Val)
	dfs(root.Right, ret)
}
