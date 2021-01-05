package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}

	ret := make([][]int, 0)
	backtrace(root, []int{}, sum, &ret, "")

	return ret
}

func backtrace(root *TreeNode, sub []int, sum int, ret *[][]int, pos string) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		if sum == root.Val {
			tmp := make([]int, len(sub))
			copy(tmp, sub)
			*ret = append(*ret, append(tmp, root.Val))
		}

		return
	}

	backtrace(root.Left, append(sub, root.Val), sum-root.Val, ret, "left")
	backtrace(root.Right, append(sub, root.Val), sum-root.Val, ret, "right")
}
