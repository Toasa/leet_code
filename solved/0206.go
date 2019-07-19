/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
**/

// 1 -> 2 -> 3 -> 4 -> 5
// 5 -> 4 -> 3 -> 2 -> 1

func newListNode(val int) *ListNode {
	return &ListNode{Val: val}
}

func rec(curNode, headRev *ListNode) *ListNode {
	if curNode == nil {
		return headRev
	}
	tmp := newListNode(curNode.Val)
	tmp.Next = headRev
	headRev = tmp
	curNode = curNode.Next
	return rec(curNode, headRev)
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	headRev := newListNode(head.Val)
	return rec(head.Next, headRev)
}