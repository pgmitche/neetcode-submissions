type MyStack struct {
	lifo []int
}

func Constructor() MyStack {
	return MyStack{lifo: []int{}}
}

func (this *MyStack) Push(x int) {
	// we need to prepend to the array
	this.lifo = append([]int{x}, this.lifo...)
}

func (this *MyStack) Pop() int {
	top := this.Top()
	// remove from FRONT of array
	this.lifo = this.lifo[1:]
	return top
}

func (this *MyStack) Top() int {
    return this.lifo[0]
}

func (this *MyStack) Empty() bool {
	return len(this.lifo) == 0
}
