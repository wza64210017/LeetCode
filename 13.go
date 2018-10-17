package main

import (
	"fmt"
)

func romanToInt(s string) int {
	m := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	sum := 0
	for i := len(s) - 1; i >= 0; i-- {
		str := string(s[i])
		if i == len(s) - 1 {
			sum += m[str]
			continue
		}

		if m[str] < m[string(s[i+1])] {
			sum -= m[str]
		} else {
			sum += m[str]
		}
	}

	return sum
}

func main() {
	fmt.Println(romanToInt("LVIII"))
}
