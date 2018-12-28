package main

import (
	"fmt"
	"unicode/utf8"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	//周期
	T := 2 * (numRows - 1)
	length := len(s)

	str := make([]byte, length)

	for i := 1; i <= numRows; i++ {
		n := 1
		for j := i - 1; j < length; j += T {
			if i == 1 || i == numRows {
				str = append(str, s[j])
			} else {
				str = append(str, s[j])
				if n*T-j < length {
					str = append(str, s[n*T-j])
				}
				n += 2
			}
		}
	}

	return string(str)
}

func main() {
	fmt.Println(utf8.RuneCountInString("㕎檕獶栲123"))
	fmt.Println(convert("LEETCODEISHIRING", 3))
}
