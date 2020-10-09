package main

import (
	"fmt"
	"strings"
)

func myAtoi(str string) int {
	max := 1 << 31
	num := 0

	str = strings.Trim(str, " ")

	if len(str) == 0 {
		return num
	}

	pre := false
	for k, v := range str {
		if k == 0 && (v == '-' || v == '+') {
			if v == '-' {
				pre = true
			}
			continue
		}

		if v >= '0' && v <= '9' {
			num = num*10 + int(v-48)
			if pre == true && num > max {
				num = max
				break
			} else if pre == false && num > max-1 {
				num = max - 1
				break
			}

			continue
		}

		break
	}

	if pre == true {
		return -num
	}

	return num
}

func main() {
	fmt.Println(myAtoi("2147483648"))
}
