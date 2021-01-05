package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) []int {
	res := make([][]int, 0)
	dfs(root, &res, 0)

	data := make([]int, 0, len(res)*2)
	for _, v := range res {
		data = append(data, v...)
	}

	return data
}

func dfs(root *TreeNode, res *[][]int, level int) {
	if root == nil {
		return
	}

	if len(*res) == level {
		*res = append(*res, []int{})
	}

	(*res)[level] = append((*res)[level], root.Val)
	dfs(root.Left, res, level+1)
	dfs(root.Right, res, level+1)
}
