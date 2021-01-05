package main

import (
	"fmt"
	"strconv"
)

var data = map[string]string{
	"0": "a", "1": "b", "2": "c", "3": "d", "4": "e", "5": "f", "6": "g", "7": "h",
	"8": "i", "9": "j", "10": "k", "11": "l", "12": "m", "13": "n", "14": "o", "15": "p",
	"16": "q", "17": "r", "18": "s", "19": "t", "20": "u", "21": "v", "22": "w", "23": "x",
	"24": "y", "25": "z",
}

func translateNum(num int) int {
	s := strconv.Itoa(num)
	ret := make(map[string]bool)

	backTrace(s, "", &ret)
	fmt.Println(ret)
	return len(ret)
}

func backTrace(s string, sub string, ret *map[string]bool) {
	if len(s) == 0 {
		(*ret)[sub] = true
		return
	}

	backTrace(s[1:], sub+data[s[:1]], ret)
	if len(s) >= 2 && s[:2] <= "25" && s[:2] >= "10" {
		backTrace(s[2:], sub+data[s[:2]], ret)
	}
}

func main() {
	//fmt.Println(translateNum(12258))
	fmt.Println(translateNum(506))
}
