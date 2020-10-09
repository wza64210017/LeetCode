package main

import "fmt"

// 约瑟夫环 https://blog.csdn.net/u011500062/article/details/72855826
//func lastRemaining(n int, m int) int {
//	if n == 1 {
//		return 0
//	}
//
//	return (lastRemaining(n-1, m) + m) % n
//}

func lastRemaining(n int, m int) int {
	remain := 0
	for i := 2; i <= n; i++ {
		remain = (remain + m) % i
		fmt.Println("i", i, "remain", remain)
	}

	return remain
}

func main() {
	fmt.Println(lastRemaining(5, 3))
}
