/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Tr struct {
	depth int
	box   [][]int
}

var tr *Tr

func traverse(node *TreeNode, depth int) {
	if depth < tr.depth {
		tr.box[depth] = append(tr.box[depth], node.Val)
	} else {
		tr.depth++
		tr.box = append(tr.box, []int{node.Val})
	}

	if node.Left != nil {
		traverse(node.Left, depth+1)
	}
	if node.Right != nil {
		traverse(node.Right, depth+1)
	}
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	tr = &Tr{depth: 0}
	traverse(root, 0)

	return tr.box
}