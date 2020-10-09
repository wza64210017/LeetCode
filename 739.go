package main

import "fmt"

func dailyTemperatures(T []int) []int {
	length := len(T)
	if length == 0 {
		return T
	}

	newT := make([]int, length)
	stack := make([]int, 0, length)
	for i := length - 1; i >= 0; i-- {
		stackLen := len(stack)
		for stackLen != 0 && T[i] >= T[stackPeek(stack)]{
			stack = stack[:stackLen-1]
			stackLen = len(stack)
		}

		newT[i] = 0
		if stackLen != 0 {
			newT[i] = stackPeek(stack) - i
		}

		stack = append(stack, i)
	}

	return newT
}

func stackPeek(stack []int) int {
	return stack[len(stack)-1]
}

func main() {
	T := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(T))
}
