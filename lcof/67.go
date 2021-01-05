package main

import (
	"fmt"
	"strings"
)

const MIN = 1 << 31

func strToInt(str string) int {
	str = strings.Trim(str, " ")
	if len(str) == 0 {
		return 0
	}

	var flag bool
	flag, str = pre(str)

	num := 0
	for _, v := range str {
		if v < '0' || v > '9' {
			break
		}

		num = num*10 + int(v-'0')
		if f, v := process(flag, num); f {
			return v
		}
	}

	_, v := process(flag, num)
	return v
}

// 判断正负
func pre(str string) (bool, string) {
	if str[0] == '-' {
		return true, str[1:]
	}

	if str[0] == '+' {
		return false, str[1:]
	}

	return false, str
}

func process(flag bool, num int) (bool, int) {
	switch {
	case flag && num < MIN:
		return false, -num
	case flag && num >= MIN:
		return true, -MIN
	case !flag && num < MIN-1:
		return false, num
	default:
		return true, MIN - 1
	}
}

func main() {
	fmt.Println(strToInt("42"))
	fmt.Println(strToInt("   -42"))
	fmt.Println(strToInt("4193 with words"))
	fmt.Println(strToInt("words and 987"))
	fmt.Println(strToInt("-91283472332"))
	fmt.Println(strToInt("+1"))
	fmt.Println(strToInt("2147483648"))
}
