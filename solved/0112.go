/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var S int

func isLeaf(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}

func calc(node *TreeNode, val int) bool {
	if isLeaf(node) {
		return val+node.Val == S
	}

	if node.Left != nil {
		if calc(node.Left, val+node.Val) {
			return true
		}
	}

	if node.Right != nil {
		if calc(node.Right, val+node.Val) {
			return true
		}
	}
	return false
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}

	S = sum
	return calc(root, 0)
}