package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	dfs(root, 0)
	return root
}

func dfs(root *TreeNode, val int) int {
	if root == nil {
		return val
	}

	root.Val += dfs(root.Right, val)
	return dfs(root.Left, root.Val)
}

func main() {

}
