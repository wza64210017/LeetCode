package main

import "fmt"

type MyQueue struct {
	List []int
	Len  int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	queue := MyQueue{
		List: make([]int, 0),
		Len: 0,
	}

	return queue
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.List = append(this.List, x)
	this.Len++
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.List) == 0 {
		return 0
	}

	value := this.List[0]
	this.List = this.List[1:]
	this.Len--

	return value
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.List) == 0 {
		return 0
	}

	value := this.List[0]
	return value
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.Len == 0
}

func main() {
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	fmt.Println(queue.Peek())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Empty())
}
