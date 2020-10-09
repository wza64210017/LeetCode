package main

import (
	"fmt"
	"strconv"
)

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	len1 := len(num1)
	len2 := len(num2)

	ret := make([]int, len1+len2)
	for i := len2 - 1; i >= 0; i-- {
		for j := len1 - 1; j >= 0; j-- {
			pos := (len1 - j - 1) + (len2 - i - 1)

			ret[pos] += multi(num2[i], num1[j])
			if ret[pos] >= 10 {
				tmp := ret[pos]
				ret[pos] = tmp % 10
				ret[pos+1] += tmp / 10
			}
		}
	}

	s := ""
	flag := false
	for i := len(ret) - 1; i >= 0; i-- {
		if ret[i] != 0 && !flag {
			flag = true
		}

		if flag {
			s += strconv.Itoa(ret[i])
		}
	}

	return s
}

func multi(a, b uint8) int {
	return int(a-48) * int(b-48)
}

func main() {
	fmt.Println(multiply("123", "456"))
}
