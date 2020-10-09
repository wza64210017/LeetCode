package main

import "fmt"

func reverseLeftWords(s string, n int) string {
	arr := make([]byte, 0, len(s))

	arr = append(arr, s[n:]...)
	arr = append(arr, s[:n]...)
	return string(arr)
}

func main() {
	s := `lrloseumgh`
	fmt.Println(reverseLeftWords(s, 6))
}
