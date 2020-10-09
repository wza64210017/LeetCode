package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	ret := make([]int, 0)
	dfs(root, &ret)

	return ret
}

func dfs(root *TreeNode, ret *[]int) {
	if root == nil {
		return
	}

	dfs(root.Left, ret)
	dfs(root.Right, ret)
	*ret = append(*ret, root.Val)
}
