package main

import "fmt"

func strStr(haystack string, needle string) int {

	if needle == "" {
		return 0
	}

	l := len(needle)
	for i := 0; i < len(haystack)-l+1; i++ {
		if string(haystack[i:l+i]) == needle {
			return i
		}
	}

	return -1
}

func main() {
	haystack := "aaaaa"
	needle := "ba"

	fmt.Println(strStr(haystack, needle))
}
