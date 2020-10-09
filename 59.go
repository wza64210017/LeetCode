package main

import (
	"fmt"
)

func generateMatrix(n int) [][]int {
	var data [4][4]int

	a, b := 0, 0
	for i := 0; i < n*n; i ++ {
		if b < n {
			data[a][b] = i + 1
			b++
		} else {

		}

		if a < n {
			data[a][b] = i +1
			a++
		}
	}

	fmt.Println(data, b)
	return nil
}

func main() {
	n := 3
	fmt.Println(generateMatrix(n))

	//4
	//1  2  3   4
	//12 13 14  5
	//11 16	15  6
	//10  9  8  7

	//5 % 4 = 1  5 / 4 = 1
	//6 % 4 = 2  6 / 4 = 1
	//7 % 4 = 3  7 / 4 = 1
	//8 % 4 = 0  8 / 4 = 2
	//9 % 4 = 1  9 / 4 = 2
	//10 % 4 = 2 10 / 4 = 2

}
