type MyStack struct {
	lifo []int
}

func Constructor() MyStack {
	return MyStack{lifo: []int{}}
}

func (this *MyStack) Push(x int) {
	this.lifo = append(this.lifo, x)
}

func (this *MyStack) Pop() int {
	var x int
	x, this.lifo = this.lifo[len(this.lifo)-1], this.lifo[:len(this.lifo)-1]
	return x
}

func (this *MyStack) Top() int {
	return this.lifo[len(this.lifo)-1]
}

func (this *MyStack) Empty() bool {
	return len(this.lifo) == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param2 := obj.Pop();
 * param3 := obj.Top();
 * param4 := obj.Empty();
 */
