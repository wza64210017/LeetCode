package main

import "fmt"

func removeKdigits(num string, k int) string {
LOOP:
	for k > 0 {
		i := 0
		for i = 0; i < len(num)-1; i++ {
			if num[i] > num[i+1] {
				num = num[:i] + num[i+1:]
				k--
				goto LOOP
			}
		}

		num = num[:len(num)-1]
		k--
	}

	num = removeZero(num)
	if len(num) == 0 {
		return "0"
	}

	return num
}

func removeZero(num string) string {
	for len(num) > 0 && num[0] == '0' {
		num = num[1:]
	}

	return num
}

func main() {
	fmt.Println(removeKdigits("1432219", 3))
	fmt.Println(removeKdigits("10200", 1))
	fmt.Println(removeKdigits("10", 2))
	fmt.Println(removeKdigits("9", 1))
	fmt.Println(removeKdigits("100", 1))
	fmt.Println(removeKdigits("1173", 2))
}
