/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSymmetricSub(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l != nil && r == nil {
		return false
	}
	if l == nil && r != nil {
		return false
	}

	if !isSymmetricSub(l.Left, r.Right) {
		return false
	}

	if !isSymmetricSub(l.Right, r.Left) {
		return false
	}

	return l.Val == r.Val
}

func isSymmetric(root *TreeNode) bool {
    if root == nil {
		return true
	}
	if root.Left != nil && root.Right == nil {
		return false
	}
	if root.Left == nil && root.Right != nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	if !isSymmetricSub(root.Left, root.Right) {
		return false
	}
	return true
}