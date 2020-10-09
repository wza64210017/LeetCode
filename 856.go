package main

import (
	"fmt"
)

func scoreOfParentheses(S string) int {

	stock := make([]int32, 0, len(S))
	for _, v := range S {
		if v == '(' {
			stock = append(stock, v)
			continue
		}

		var tmp int32 = 0
		for i := len(stock) - 1; i >= 0; i-- {
			if stock[i] != '(' {
				tmp += stock[i]
				continue
			}

			stock[i] = 1
			if tmp != 0 {
				stock[i] = 2 * tmp
			}

			stock = stock[:i+1]
			break
		}
	}

	var sum int32
	for _, v := range stock {
		sum += v
	}

	return int(sum)
}

func main() {
	S := "(()(()))"
	//()
	fmt.Println(scoreOfParentheses(S))
}
