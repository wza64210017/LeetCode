package main

import "fmt"

func lengthOfLastWord(s string) int {
	if s == "" {
		return 0
	}

	count := 0
	pos := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != 32 {
			pos = i
			break
		}
	}

	for _, v := range s[:pos+1] {
		if v == 32 {
			count = 0
			continue
		}

		count++
	}

	return count
}

func main() {
	fmt.Println(lengthOfLastWord(""))
}
