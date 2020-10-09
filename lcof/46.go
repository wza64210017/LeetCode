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
	//m := make(map[string]bool)

	s := strconv.Itoa(num)
	fmt.Println(dfs(s))

	return 0
}

func dfs(s string) string {
	if len(s) < 2 {
		return data[s]
	}

	var res string
	res += data[s[:1]] + dfs(s[1:])
	if s[:2] < "25" && s[:2] > "10" {
		res += data[s[:2]] + dfs(s[2:])
	} else {
	}

	return res
}

func main() {
	fmt.Println(translateNum(12258))
}
