package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	ret := make([][]int, 0)
	dfs(root, 0, &ret)

	for i, vals := range ret {
		if i%2 == 0 {
			ret[i] = reverse(vals)
		}
	}

	return ret
}

func dfs(root *TreeNode, level int, ret *[][]int) {
	if root == nil {
		return
	}

	if len(*ret) == level {
		*ret = append(*ret, []int{})
	}

	dfs(root.Left, level+1, ret)
	dfs(root.Right, level+1, ret)
	(*ret)[level] = append((*ret)[level], root.Val)
}

func reverse(arr []int) []int {
	l := len(arr)
	for i := 0; i < l/2; i++ {
		arr[i], arr[l-i-1] = arr[l-i-1], arr[i]
	}

	return arr
}
