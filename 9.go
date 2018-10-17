package main

import (
	"strconv"
	"fmt"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	str := strconv.Itoa(x)
	length := len(str)
	for i := 0; i < length/2; i++ {
		if str[i] != str[length-1-i] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isPalindrome(10))
}
