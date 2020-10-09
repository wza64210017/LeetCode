package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	list := strings.Split(path, "/")
	arr := make([]string, 0, len(list))
	pos := 0
	for _, v := range list {
		switch v {
		case "":
			continue
		case ".":
			continue
		case "..":
			if pos > 0 {
				arr = arr[:pos-1]
				pos--
			}
		default:
			arr = append(arr, v)
			pos++
		}
	}

	return "/" + strings.Join(arr, "/")
}

func main() {
	str := "/home/"
	fmt.Println(simplifyPath(str))
}
