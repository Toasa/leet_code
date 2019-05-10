// Maximum Depth of Binary Tree

var max_d int

func calc(node *TreeNode, d int) {
	if node.Left == nil && node.Right == nil {
		if max_d < d {
			max_d = d
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

func maxDepth(root *TreeNode) int {
	max_d = 0
	if root == nil {
		return 0
	}

	calc(root, 1)
	return max_d
}