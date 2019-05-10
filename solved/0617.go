/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func newNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

func merge(t1, t2 *TreeNode) *TreeNode {
	node := newNode(t1.Val + t2.Val)

	if t1.Left == nil && t2.Left == nil {
		node.Left = nil
	} else if t1.Left != nil && t2.Left == nil {
		node.Left = t1.Left
	} else if t1.Left == nil && t2.Left != nil {
		node.Left = t2.Left
	} else {
		node.Left = merge(t1.Left, t2.Left)
	}

	if t1.Right == nil && t2.Right == nil {
		node.Right = nil
	} else if t1.Right != nil && t2.Right == nil {
		node.Right = t1.Right
	} else if t1.Right == nil && t2.Right != nil {
		node.Right = t2.Right
	} else {
		node.Right = merge(t1.Right, t2.Right)
	}

	return node
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}

	return merge(t1, t2)
}