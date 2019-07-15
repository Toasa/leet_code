/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func detectCycle(head *ListNode) *ListNode {
	p1 := head
	p2 := head
	var err1 error
	var err2 error
	
	for p1 != nil || p2 != nil {
		p1, err1 = move(p1)
		p2, err2 = moveTwice(p2)

		if err1 != nil || err2 != nil {
			return nil
		}
		if p1 == p2 {
			return loopStartPoint(head)
		}
	}
	return nil
}

func move(p *ListNode) (*ListNode, error) {
	if p.Next == nil {
		return nil, errors.New("next node is nil")
	}
	return p.Next, nil
}

func moveTwice(p *ListNode) (*ListNode, error) {
	if p.Next == nil {
		return nil, errors.New("next node is nil")
	}
	if p.Next.Next == nil {
		return nil, errors.New("next next node is nil")
	}
	return p.Next.Next, nil
}

func loopStartPoint(p *ListNode) *ListNode {
	m := make(map[*ListNode]bool)
	m[p] = true

	for {
		p = p.Next
		if m[p] {
			return p
		}
		m[p] = true
	}
	return nil
}