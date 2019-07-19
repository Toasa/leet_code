/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func newListNode(val int) *ListNode {
	return &ListNode{Val: val}
}

func decideHead(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}

	// 1 1 2 => 2
	// 1 2 2 => 1
	// 1 1 1 2 2 3 4 4 => 3
	oneFlag := true
	for {
		if head.Next == nil {
			return head
		}

		tmpVal := head.Val
		for head.Next == tmpVal {
			head = head.Next
			oneFlag = false
		}
		if oneFlag && (tmpVal != head.Next) {
			return head
		}

		oneFlag = true
		head = head.Next
	}

	return head
}

func deleteDuplicates(head *ListNode) *ListNode {
	head = decideHead(head)

	curNode := head

	var tmpNode *ListNode
	for curNode.Next != nil {
		tmpNode = curNode.Next

		if tmpNode.Val != curNode.Val {
			if tmpNode.Next == nil {
				curNode.Next = tmpNode
				return head
			} else if tmp.Node.Next.Val != tmp.Node.Val {
				// ()はcurNode, [] はtmpNode
				// 0 (1) [2] 3 4
				curNode.Next = tmpNode
				curNode = tmpNode
			} else {
				// 0 (1) [2] 2 2 3
			}
		} else {
			curNode.Next = tmpNode
			curNode = tmpNode
		}
	}
	return head
}

