package main

import "fmt"

func isHappy(n int) bool {
	m := make(map[int]bool)

	for {
		val := cal(subs(n))
		if val == 1 {
			return true
		}

		if m[val] {
			break
		}

		m[val] = true
		n = val
	}

	return false
}

func subs(num int) []int {
	ret := make([]int, 0)
	for num != 0 {
		ret = append(ret, num%10)
		num /= 10
	}

	return ret
}

func cal(nums []int) int {
	val := 0
	for _, num := range nums {
		val += num * num
	}

	return val
}

func main() {
	fmt.Println(isHappy(100))
}
