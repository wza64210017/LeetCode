package main

import "fmt"

func minAddToMakeValid(S string) int {
	if len(S) == 0 {
		return 0
	}

	l := 0
	r := 0
	for _, v := range S {
		if v == '(' {
			l++
		} else if v == ')' {
			if l == 0 {
				r++
			} else {
				l--
			}
		}
	}

	return l + r
}

func main() {
	S := "()))(("
	fmt.Println(minAddToMakeValid(S))
}
