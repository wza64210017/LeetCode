package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Node struct {
	Count float64
	Sum   float64
}

func averageOfLevels(root *TreeNode) []float64 {
	nodes := make([]Node, 0)
	travel(root, 0, &nodes)

	ret := make([]float64, 0, len(nodes))
	for _, node := range nodes {
		ret = append(ret, node.Sum/node.Count)
	}

	return ret
}

func travel(root *TreeNode, level int, nodes *[]Node) {
	if root == nil {
		return
	}

	if len(*nodes) == level {
		*nodes = append(*nodes, Node{})
	}

	node := (*nodes)[level]
	node.Count += 1
	node.Sum += float64(root.Val)
	(*nodes)[level] = node

	travel(root.Left, level+1, nodes)
	travel(root.Right, level+1, nodes)
}

func main() {

}
