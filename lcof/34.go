package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	ret := make([][]int, 0)
	backtrace(root, []int{}, sum, &ret)

	return ret
}

func backtrace(root *TreeNode, sub []int, sum int, ret *[][]int) {
	if root == nil || sum < 0 {
		return
	}

	if sum == 0 {
		tmp := make([]int, len(sub))
		copy(tmp, sub)
		*ret = append(*ret, tmp)
		return
	}

	backtrace(root.Left, append(sub, root.Val), sum-root.Val, ret)
	backtrace(root.Right, append(sub, root.Val), sum-root.Val, ret)
}
