/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	cur := head
	var l int
	for l = 1; cur.Next != nil; l++ {
		cur = cur.Next
	}

	if l == 1 {
		return nil
	}
	if l == 2 {
		if n == 1 {
			head.Next = nil
			return head
		}
		if n == 2 {
			return head.Next
		}
	}

	// l >= 3
	if n == l {
		return head.Next
	}

	cur = head
	for i := 1; i < l-n; i++ {
		cur = cur.Next
	}

	cur.Next = cur.Next.Next

	return head
}