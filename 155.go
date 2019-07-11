package main

import (
	"fmt"
)

type MinStack struct {
	List []int
	Len  int
	Min  int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	stack := MinStack{
		List: make([]int, 0),
		Len:  0,
		Min:  0,
	}

	return stack
}

func (this *MinStack) Push(x int) {
	this.List = append(this.List, x)
	this.Len++

	if this.Min > x || this.Len == 1 {
		this.Min = x
	}
}

func (this *MinStack) Pop() {
	if this.Len == 0 {
		return
	}

	needFlush := false
	if this.Min == this.Top() {
		needFlush = true
	}

	this.List = this.List[:this.Len-1]
	this.Len--

	if needFlush == true {
		min := this.Top()
		for _, value := range this.List {
			if value < min {
				min = value
			}
		}

		this.Min = min
	}
}

func (this *MinStack) Top() int {
	if this.Len == 0 {
		return 0
	}

	return this.List[this.Len-1]
}

func (this *MinStack) GetMin() int {
	return this.Min
}

func main() {
	stack := Constructor()
	stack.Push(2)
	stack.Push(0)
	stack.Push(3)
	stack.Push(0)
	fmt.Println(stack.GetMin())
	stack.Pop()
	fmt.Println(stack.GetMin())
	stack.Pop()
	fmt.Println(stack.GetMin())
	stack.Pop()
	fmt.Println(stack.GetMin())
	//
	//stack := Constructor()
	//stack.Push(1)
	//stack.Push(2)
	//fmt.Println(stack.Top())
	//fmt.Println(stack.GetMin())
	//stack.Pop()
	//fmt.Println(stack.GetMin())
	//fmt.Println(stack.Top())
}
