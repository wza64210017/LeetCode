package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	ret := make([]string, 0)
	for i := 0; i < len(path); i++ {
		tmp := ""
		for i < len(path) && path[i] != '/' {
			tmp += string(path[i])
			i++
		}

		if tmp == ".." && len(ret) > 0 {
			ret = ret[:len(ret)-1]
		}

		if tmp != "." && tmp != "" && tmp != ".." {
			ret = append(ret, tmp)
		}
	}

	return "/" + strings.Join(ret, "/")
}

func main() {
	fmt.Println(simplifyPath("/a//b////c/d//././/.."))
}
