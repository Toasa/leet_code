type MinStack struct {
	stk []int
	min int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack {
		min: 2147483647,
	}
}


func (this *MinStack) Push(x int)  {
    if x < this.min {
		this.min = x
	}
	this.stk = append(this.stk, x)
}


func (this *MinStack) Pop()  {
	top := this.Top()
	this.stk = this.stk[:len(this.stk) - 1]
	if this.min == top {
		min := 2147483647
		for i := 0; i < len(this.stk); i++ {
			if this.stk[i] < min {
				min = this.stk[i]
			}
		}
		this.min = min
	}    
}


func (this *MinStack) Top() int {
    return this.stk[len(this.stk) - 1]
}


func (this *MinStack) GetMin() int {
    return this.min
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */