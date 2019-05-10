import (
	"sort"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var arr []int

func newNode(val int) *TreeNode {
	return &TreeNode{Val: val}
}

func makeBST(min, max int) *TreeNode {
	// -10 -3 0 5 9
	// 0   1  2 3 4

	// makeBST(2, 4)
	// min = 2, mid = 3, max = 4

	// makeBST(2, 3)
	// min = 2, mid = 2, max = 3

	// makeBST(3, 4)
	// min = 3, mid = 3, max = 4

	mid := (min + max) / 2
	node := newNode(arr[mid])

	if mid-min > 1 {
		node.Left = makeBST(min, mid)
	} else if mid-min == 1 {
		node.Left = newNode(arr[min])
		node.Right = nil
		return node
	} else if mid-min == 0 {
		return nil
	}

	if max-mid > 1 {
		node.Right = makeBST(mid, max)
	} else if max-mid == 1 {
		// ここの処理が走っていない
		node.Left = nil
		node.Right = newNode(arr[max])
		return node
	}

	return node
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}

	sort.Ints(nums)
	arr = nums

	root := makeBST(0, len(arr)-1)
	return root
}