package main

import (
	"fmt"
)

func generateParenthesis(n int) []string {
	ret := make([]string, 0)
	backtrace(n, 0, 0, "", &ret)

	return ret
}

func backtrace(n, l, r int, sub string, ret *[]string) {
	// 左括号数 大于 n，或者 左括号数小于右括号数
	if l > n || l < r {
		return
	}

	if l+r == n*2 {
		*ret = append(*ret, sub)
		return
	}

	backtrace(n, l+1, r, sub+"(", ret)
	backtrace(n, l, r+1, sub+")", ret)
}

func main() {
	fmt.Println(generateParenthesis(3))
}
