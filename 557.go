package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	arr := strings.Split(s, " ")
	for k, v := range arr {
		arr[k] = reverse1(v)
	}

	return  strings.Join(arr, " ")
}

func reverse1(s string) string {
	length := len(s)
	if length == 1 {
		return s
	}

	sbyte := make([]byte, length)
	for i := 0; i < (length+1)/2; i++ {
		sbyte[i], sbyte[length-1-i] = s[length-i-1], s[i]
	}

	return string(sbyte)
}

func main() {
	s := "Let's take LeetCode contest"
	//s = "I love u"
	fmt.Println(reverseWords(s))
}
