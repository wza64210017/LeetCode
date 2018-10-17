package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	s := ""
	j := 0
	for {
		var tmp uint8 = 0
		for _, str := range strs {
			if j > (len(str) - 1) {
				return s
			}
			if tmp == 0 {
				tmp = str[j]
				continue
			}

			if str[j] != tmp {
				return s
			}
		}
		s += string(strs[0][j])
		j++
	}

	return s
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"dog"}))
}
