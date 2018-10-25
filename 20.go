package main

import "fmt"

func isValid(s string) bool {

	m := map[string]string{")": "(", "]": "[", "}": "{"}
	arr := []string{}
	for _, v := range s {
		str := string(v)
		if str == "{" || str == "[" || str == "(" {
			arr = append(arr, str)
		} else {
			length := len(arr)
			if length == 0 || arr[length-1] != m[str] {
				return false
			}

			if arr[length-1] == m[str] {
				arr = arr[:length-1]
			}
		}
	}

	if len(arr) == 0 {

		return true
	}
	return false
}

func main() {
	fmt.Println(isValid("(])"))
}
