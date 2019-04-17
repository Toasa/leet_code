/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil && l2 == nil {
		return nil
	} else if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	var head *ListNode
	var cur *ListNode

	var l1_tail bool = false
	var l2_tail bool = false

	if l1.Val < l2.Val {
		head = l1
		cur = l1
		if l1.Next != nil {
			l1 = l1.Next
		} else {
			l1_tail = true
		}
	} else {
		head = l2
		cur = l2
		if l2.Next != nil {
			l2 = l2.Next
		} else {
			l2_tail = true
		}
	}

	for !l1_tail && !l2_tail {
		if l1.Val < l2.Val {
			cur.Next = l1
			cur = cur.Next
			if l1.Next != nil {
				l1 = l1.Next
			} else {
				l1_tail = true
			}
		} else {
			cur.Next = l2
			cur = cur.Next
			if l2.Next != nil {
				l2 = l2.Next
			} else {
				l2_tail = true
			}
		}
	}

	if l1_tail {
		cur.Next = l2
	} else {
		cur.Next = l1
	}
	return head
}