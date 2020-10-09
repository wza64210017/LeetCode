package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	s = strings.Trim(s, " ")

	pos := 0
	ret := make([]string, 0)
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			ret = append(ret, s[pos:i])
			for s[i] == ' ' {
				i++
			}

			pos = i
			continue
		}
	}

	ret = append(ret, s[pos:])

	str := ""
	for i := len(ret) - 1; i >= 0; i-- {
		str += ret[i] + " "
	}

	return str[:len(str)-1]
}

func main() {
	fmt.Println(reverseWords("the sky is blue"))
	fmt.Println(reverseWords("  hello world!  "))
	fmt.Println("+" + reverseWords("a good   example") + "+")
}
