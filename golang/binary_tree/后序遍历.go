package binarytree

func postorderTraversal(root *TreeNode) []int {
	var res []int
	var traversal func(*TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		traversal(node.Right)
		res = append(res, node.Val)
	}

	traversal(root)
	return res
}
