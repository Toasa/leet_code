type Node struct {
	val int
	next *Node
}

type KthLargest struct {
	K int
	head *Node
	elemNum int
}

func makeLinkedList(nums []int) *Node {
	if len(nums) == 0 {
		return nil
	}

	// numsをソートする
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	head := &Node{val: nums[0]}
	cur := head

	for i := 1; i < len(nums); i++ {
		tmp := &Node{val: nums[i], next: nil}
		cur.next = tmp
		cur = tmp
	}
	return head
}

func (this *KthLargest) appendList(newVal int) {
	newNode := &Node{val: newVal, next: nil}
	if this.head == nil {
		this.head = newNode
		this.elemNum++
		return
	}

	// 先頭に挿入
	if this.head.val <= newVal {
		newNode.next = this.head
		this.head = newNode
		this.elemNum++
		return
	}

	// 中間に挿入
	pivot := this.head
	for i := 0; i < this.elemNum - 1; i++ {
		if pivot.next.val <= newVal && newVal <= pivot.val {
			newNode.next = pivot.next
			pivot.next = newNode
			this.elemNum++
			return
		}
		pivot = pivot.next
	}

	// 末尾に挿入
	pivot.next = newNode	

	this.elemNum++
}

func Constructor(k int, nums []int) KthLargest {
	head := makeLinkedList(nums)
	KL := KthLargest{K: k, head: head, elemNum: len(nums)}
	return KL
}

func (this *KthLargest) printList() {
	for node := this.head; node != nil; node = node.next {
		fmt.Printf(" %d ->", node.val)
	}
	fmt.Println()
}

func (this *KthLargest) Add(val int) int {
	this.appendList(val)
	
	node := this.head
	for i := 1; i < this.K; i++ {
		node = node.next
	}
	return node.val
}


/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */