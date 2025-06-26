package binarytree

func isSymmetric(root *TreeNode) bool {
	var defs func(left, right *TreeNode) bool
	defs = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}
		return defs(left.Left, right.Right) && defs(right.Left, left.Right)
	}

	return defs(root.Left, root.Right)
}
