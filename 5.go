package main

import "fmt"

func longestPalindrome(s string) string {

	length := len(s)
	palindrome := make([][]bool, length)
	maxLength := 0
	longestStr := ""

	for i := 0; i < length; i++ {
		j := i
		palindrome[j] = make([]bool, length)

		for j >= 0 {
			if s[i] == s[j] && (i-j < 2 || palindrome[j+1][i-1] == true) {
				palindrome[j][i] = true

				if i-j+1 >= maxLength {
					longestStr = s[j : i+1]
					maxLength = i - j + 1
				}
			}
			j--
		}
	}

	return longestStr
}

func main() {
	fmt.Println(longestPalindrome("abcda"))
	//fmt.Println(longestPalindrome("ccc"))
	fmt.Println(longestPalindrome("babad"))
}
