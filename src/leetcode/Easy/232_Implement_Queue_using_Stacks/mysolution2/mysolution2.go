package main

import (
	"github.com/emirpasic/gods/stacks/linkedliststack"
)

type MyQueue struct {
	in  *linkedliststack.Stack
	out *linkedliststack.Stack
}

func Constructor() MyQueue {
	return MyQueue{
		in:  linkedliststack.New(),
		out: linkedliststack.New(),
	}
}

func (this *MyQueue) Push(x int) {
	this.in.Push(x)
}

func (this *MyQueue) Pop() int {
	if this.out.Empty() {
		this.fromStackToQueue()
	}
	res, _ := this.out.Pop()
	return res.(int)
}

func (this *MyQueue) Peek() int {
	if this.out.Empty() {
		this.fromStackToQueue()
	}
	res, _ := this.out.Peek()
	return res.(int)
}

func (this *MyQueue) Empty() bool {
	if this.out.Empty() {
		this.fromStackToQueue()
	}
	return this.out.Empty()
}

func (this *MyQueue) fromStackToQueue() {
	for {
		p, ok := this.in.Pop()
		if !ok {
			break
		}
		this.out.Push(p)
	}
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
