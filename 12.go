package main

import "fmt"

func intToRoman(num int) string {
	var M = map[int]string{
		1:    "I",
		5:    "V",
		10:   "X",
		50:   "L",
		100:  "C",
		500:  "D",
		1000: "M",
	}

	s := ""
	loop := 1

	for num > 0 {
		left := num % 10
		if left <= 3 {
			tmp := ""
			for left > 0 {
				tmp += M[loop]
				left--
			}
			s = tmp + s
		} else if left == 4 {
			s = M[loop] + M[loop*5] + s
		} else if left == 5 {
			s = M[loop*5] + s
		} else if left <= 8 {
			tmp := M[loop*5]
			for left > 5 {
				tmp += M[loop]
				left--
			}
			s = tmp + s
		} else if left == 9 {
			s = M[loop] + M[loop*10] + s
		}

		num = num / 10
		loop *= 10
	}

	return s
}

func main() {
	fmt.Println(intToRoman(1994))
	fmt.Println(intToRoman(58))
	fmt.Println(intToRoman(9))
	fmt.Println(intToRoman(4))
	fmt.Println(intToRoman(3))
	fmt.Println(intToRoman(10))
	fmt.Println(intToRoman(3999))
	fmt.Println(intToRoman(101))
}
