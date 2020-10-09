package main

type MaxQueue struct {
	List []int
	Max  int
	Size int
}

func Constructor() MaxQueue {
	return MaxQueue{
		make([]int, 0),
		0,
		0,
	}
}

func (this *MaxQueue) Max_value() int {
	if this.Size == 0 {
		return -1
	}

	return this.Max
}

func (this *MaxQueue) Push_back(value int) {
	this.List = append(this.List, value)
	if this.Max < value {
		this.Max = value
	}

	this.Size++
}

func (this *MaxQueue) Pop_front() int {
	if this.Size == 0 {
		return -1
	}

	val := this.List[0]
	this.List = this.List[1:]
	if val == this.Max {
		// 重新计算最大值
		max := 0
		for _, v := range this.List {
			if v > max {
				max = v
			}
		}

		this.Max = max
	}

	this.Size--
	return val
}

func main() {

}
