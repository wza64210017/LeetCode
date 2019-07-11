package main

import "fmt"

func backspaceCompare(S string, T string) bool {
	return processStr(S) == processStr(T)
}

func processStr(str string) string {
	arr := make([]int32, 0, len(str))
	for _, v := range str {
		if v == '#' {
			length := len(arr)
			if length >= 1 {
				arr = arr[:length-1]
			}
		} else {
			arr = append(arr, v)
		}
	}

	return string(arr)
}

func main() {
	S := "ab##"
	T := "c#d#"
	fmt.Println(backspaceCompare(S, T))
}
