package binarytree

func PreOrderTraversal(root *TreeNode) []int {
	res := []int{}
	var traversal func(*TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}

		res = append(res, node.Val)
		traversal(node.Left)
		traversal(node.Right)
	}

	traversal(root)
	return res
}
