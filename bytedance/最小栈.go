package main

type MinStack struct {
	list []int
	size int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		make([]int, 0),
		0,
	}
}

func (this *MinStack) Push(x int) {
	this.list = append(this.list, x)
	this.size++
}

func (this *MinStack) Pop() {
	if this.size == 0 {
		return
	}

	this.list = this.list[:this.size-1]
	this.size--
}

func (this *MinStack) Top() int {
	if this.size == 0 {
		return -1
	}

	return this.list[this.size-1]
}

func (this *MinStack) GetMin() int {
	if this.size == 0 {
		return -1
	}

	min := this.list[this.size-1]
	for _, v := range this.list {
		if min > v {
			min = v
		}
	}

	return min
}
