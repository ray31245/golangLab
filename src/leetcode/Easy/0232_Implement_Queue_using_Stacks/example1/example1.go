package main

type MyQueue struct {
	Stack *[]int
	Queue *[]int
}

func Constructor() MyQueue {
	tmp1, tmp2 := []int{}, []int{}
	return MyQueue{Stack: &tmp1, Queue: &tmp2}
}

func (this *MyQueue) Push(x int) {
	*this.Stack = append(*this.Stack, x)
}

func (this *MyQueue) Pop() int {
	if len(*this.Queue) == 0 {
		this.fromStackToQueue(this.Stack, this.Queue)
	}

	popped := (*this.Queue)[len(*this.Queue)-1]
	*this.Queue = (*this.Queue)[:len(*this.Queue)-1]
	return popped
}

func (this *MyQueue) Peek() int {
	if len(*this.Queue) == 0 {
		this.fromStackToQueue(this.Stack, this.Queue)
	}

	return (*&*this.Queue)[len(*this.Queue)-1]
}

func (this *MyQueue) Empty() bool {
	return len(*this.Queue)+len(*this.Stack) == 0
}

func (this *MyQueue) fromStackToQueue(s, q *[]int) {
	for len(*s) > 0 {
		popped := (*s)[len(*s)-1]
		*s = (*s)[:len(*s)-1]

		*q = append(*q, popped)
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
