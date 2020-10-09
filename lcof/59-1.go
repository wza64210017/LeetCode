package main

import "fmt"

type stack []int

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}

	queue := make(stack, 0, k)
	ret := make(stack, 0, len(nums)-k+1)

	for i := range nums {
		// 删除队列内小于 num[i] 的元素
		for i > 0 && queue.len() > 0 && nums[i] > queue.right() {
			queue.popRight()
		}

		// 添加 nums[i]
		queue.push(nums[i])
		// 左队首元素 等于 nums[i-k], 则移除
		if i >= k && nums[i-k] == queue[0] {
			queue.popLeft()
		}

		if i >= k-1 {
			ret = append(ret, queue[0])
		}
	}

	return ret
}

func (s *stack) len() int {
	return len(*s)
}

func (s *stack) right() int {
	return (*s)[s.len()-1]
}

func (s *stack) popRight() {
	*s = (*s)[:s.len()-1]
}

func (s *stack) popLeft() {
	*s = (*s)[1:]
}

func (s *stack) push(i int) {
	*s = append(*s, i)
}


func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	//fmt.Println(maxSlidingWindow([]int{1, -1}, 1))
}
