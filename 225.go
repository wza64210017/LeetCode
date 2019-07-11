package main

import "fmt"

type MyStack struct {
	List []int
	Len  int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	stack := MyStack{
		List: make([]int, 0),
		Len:  0,
	}

	return stack
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.List = append(this.List, x)
	this.Len++
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if this.Empty() {
		return 0
	}

	value := this.List[this.Len-1]
	this.List = this.List[:this.Len-1]
	this.Len--

	return value
}

/** Get the top element. */
func (this *MyStack) Top() int {
	if this.Empty() {
		return 0
	}

	return this.List[this.Len-1]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.Len == 0
}

func main() {
	stack := Constructor()
	stack.Push(1)
	stack.Push(2)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Top())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Empty())
}
