package main

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
	s := ""
	for n > 0 {
		s = report(s)
		n--
	}

	return s
}

func report(str string) string {
	if str == "" {
		return "1"
	}

	s := ""
	count := 1
	tmp := str[0]
	for i := 1; i < len(str); i++ {
		if str[i] == tmp {
			count++
		} else {
			s = s + strconv.Itoa(count) + string(tmp)
			tmp = str[i]
			count = 1
		}
	}

	s = s + strconv.Itoa(count) + string(str[len(str) - 1])
	return s
}

func main() {
	fmt.Println(countAndSay(6))
}

//1
//11
//21
//1211
