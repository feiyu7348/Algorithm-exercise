package binarytree

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	depth := 0

	var order func(node *TreeNode, depth int)
	order = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}

		if len(res) == depth {
			res = append(res, []int{})
		}

		res[depth] = append(res[depth], node.Val)
		order(node.Left, depth+1)
		order(node.Right, depth+1)
	}

	order(root, depth)
	return res
}
