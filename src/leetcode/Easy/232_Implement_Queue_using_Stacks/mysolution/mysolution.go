package main

type Ls struct {
	val  int
	Next *Ls
}

type MyQueue struct {
	Back  *Ls
	Front *Ls
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	if this.Back != nil {
		this.Back.Next = &Ls{x, nil}
		this.Back = this.Back.Next
	} else {
		this.Back = &Ls{x, nil}
	}
	if this.Front == nil {
		this.Front = this.Back
	}
}

func (this *MyQueue) Pop() int {
	res := this.Front.val
	this.Front = this.Front.Next
	return res
}

func (this *MyQueue) Peek() int {
	return this.Front.val
}

func (this *MyQueue) Empty() bool {
	return this.Front == nil
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
