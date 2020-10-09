package main

import (
	"fmt"
)

func removeOuterParentheses(S string) string {
	list := make([]int32, 0, 10)
	pos := 0
	for _, v := range S {
		if v == 40 {
			if pos > 0 {
				list = append(list, v)
			}
			pos++
		} else {
			if pos > 1 {
				list = append(list, v)
			}
			pos--
		}
	}

	return string(list)
}

func main() {
	str := "(()())(())"
	fmt.Println(removeOuterParentheses(str))
}
