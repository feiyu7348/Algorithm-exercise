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

/*
       1
      / \
     2   2
    / \ / \
   3  4 4  3
调用栈分析：

一、defs(2, 2)（根节点的左右子树）
	left=2, right=2（均非 nil）
	2.Val == 2.Val → 继续递归
	计算：
		defs(3, 3)（左左 vs 右右）
		defs(4, 4)（左右 vs 右左）
二、defs(3, 3)
	3.Val == 3.Val → 继续递归
	defs(nil, nil)（左左的左 vs 右右的右）→ true
	defs(nil, nil)（左左的右 vs 右右的左）→ true
	返回 true && true → true
三、defs(4, 4)
	4.Val == 4.Val → 继续递归
	defs(nil, nil)（左右的左 vs 右右的右）→ true
	defs(nil, nil)（左右的右 vs 右右的左）→ true
	返回 true && true → true
四、最终 defs(2, 2) 返回 true && true → true（树对称）。
*/
