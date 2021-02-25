package main

import (
	"fmt"
)

func addStrings(num1 string, num2 string) string {
	i, j := len(num1)-1, len(num2)-1

	arr := make([]byte, 0, i)
	var flag byte

	for i >= 0 && j >= 0 {
		n := num1[i] - '0' + num2[j] - '0' + flag
		flag = 0

		for n > 9 {
			n -= 10
			flag++
		}

		arr = append(arr, n+'0')
		i--
		j--
	}

	for i >= 0 {
		n := num1[i] - '0' + flag
		flag = 0

		for n > 9 {
			n -= 10
			flag++
		}

		arr = append(arr, n+'0')
		i--
	}

	for j >= 0 {
		n := num2[j] - '0' + flag
		flag = 0

		for n > 9 {
			n -= 10
			flag++
		}

		arr = append(arr, n+'0')
		j--
	}

	if flag > 0 {
		arr = append(arr, flag+'0')
	}

	for i := 0; i < len(arr)/2; i++ {
		arr[len(arr)-i-1], arr[i] = arr[i], arr[len(arr)-i-1]
	}

	return string(arr)
}

func main() {
	fmt.Println(addStrings("9", "99"))
}
