package main

import "fmt"

func removeDuplicates(S string) string {
	if len(S) == 1 {
		return S
	}

	list := make([]int32, 0, 100)
	for _, s := range S {
		pos := len(list) - 1
		if pos >= 0 && list[pos] == s {
			list = list[:pos]
		} else {
			list = append(list, s)
		}
	}

	return string(list)
}

func main() {
	fmt.Println(removeDuplicates("abbaca"))
}
