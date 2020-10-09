package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var num, res int
func kthLargest(root *TreeNode, k int) int {
	num = k
	dfs(root)

	return res
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}

	dfs(root.Right)
	num--
	if num == 0 {
		res = root.Val
		return
	}
	dfs(root.Left)
}