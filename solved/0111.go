import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var min_d int

func calc(node *TreeNode, d int) {
	if node.Left == nil && node.Right == nil {
		if d < min_d {
			min_d = d
		}
		return
	}

	if node.Left != nil {
		calc(node.Left, d+1)
	}
	if node.Right != nil {
		calc(node.Right, d+1)
	}
}

func minDepth(root *TreeNode) int {
	min_d = math.MaxInt32
	if root == nil {
		return 0
	}
	calc(root, 1)

	return min_d
}