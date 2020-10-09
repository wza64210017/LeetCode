package main

// 二叉树的深度

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var deep int

func maxDepth(root *TreeNode) int {
	dfs(root, 1)
	return deep
}

func dfs(root *TreeNode, level int) {
	if root == nil {
		return
	}

	if deep < level {
		deep = level
	}

	dfs(root.Left, level+1)
	dfs(root.Right, level+1)
}
