package main

import (
	"fmt"
)

//func multiply(num1 string, num2 string) string {
//	if num1 == "0" || num2 == "0" {
//		return "0"
//	}
//	len1 := len(num1)
//	len2 := len(num2)
//
//	arr1 := make([]int, 0, len1)
//	arr2 := make([]int, 0, len2)
//	for i := len1 - 1; i >= 0; i-- {
//		value, _ := strconv.Atoi(string(num1[i]))
//		arr1 = append(arr1, value)
//	}
//
//	for i := len2 - 1; i >= 0; i-- {
//		value, _ := strconv.Atoi(string(num2[i]))
//		arr2 = append(arr2, value)
//	}
//
//	ret := make([]int, len1+len2)
//	for i, a := range arr1 {
//		for k, b := range arr2 {
//			ret[i+k] += a * b
//			if ret[i+k] >= 10 {
//				tmp := ret[i+k]
//				ret[i+k] = tmp % 10
//				ret[i+k+1] += tmp / 10
//			}
//		}
//	}
//
//	pos := 0
//	for i := len(ret); i >= 0; i-- {
//		if ret[i-1] != 0 {
//			pos = i
//			break
//		}
//	}
//
//	ret = ret[:pos]
//	bytes := make([]byte, 0, len(ret))
//	for i := len(ret) - 1; i >= 0; i-- {
//		bytes = append(bytes, byte(ret[i]) + 48)
//	}
//
//	return string(bytes)
//}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	len1 := len(num1)
	len2 := len(num2)

	ret := make([]int, len1+len2)
	for i := len1 - 1; i >= 0; i-- {
		for j := len2 - 1; j >= 0; j-- {
			pos := (len1 - 1 - i) + (len2 - 1 - j)
			ret[pos] += multi(num1[i], num2[j])

			if ret[pos] >= 10 {
				tmp := ret[pos]
				ret[pos] = tmp % 10
				ret[pos+1] += tmp / 10
			}
		}
	}

	flag := false
	str := ""

	for i := len1 + len2 - 1; i >= 0; i-- {
		if ret[i] != 0 {
			flag = true
		}

		if flag == true {
			str += fmt.Sprintf("%d", ret[i])
		}
	}

	return str
}

func multi(a, b uint8) int {
	x := int(a - 48)
	y := int(b - 48)

	m := x * y
	return m
}

func main() {
	num1 := "123"
	num2 := "456"
	//1 2 3
	//    5

	//  1 5      5 1 0 0
	//1 0        5 1 1 0
	//5          5 1 6 0
	fmt.Println(multiply(num1, num2))
}
