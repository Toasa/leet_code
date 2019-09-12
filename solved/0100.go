/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil && q != nil {
		return false
	}
	if p != nil && q == nil {
		return false
	}

	if p.Left == nil && q.Left != nil {
		return false
	}
	if p.Left != nil && q.Left == nil {
		return false
	}
	if p.Right == nil && q.Right != nil {
		return false
	}
	if p.Right != nil && q.Right == nil {
		return false
	}

    if p.Val != q.Val {
		return false
	}

	if !isSameTree(p.Left, q.Left) {
		return false
	}
	if !isSameTree(p.Right, q.Right) {
		return false
	}

	return true
}