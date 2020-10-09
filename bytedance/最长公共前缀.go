package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	s := ""
	i := 0
	for {
		if len(strs[0]) <= i {
			break
		}

		tmp := strs[0][i]
		for _, str := range strs {
			if i > len(str)-1 {
				return s
			}

			if str[i] != tmp {
				return s
			}
		}

		s += string(tmp)
		i++
	}

	return s
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"dog", "doe"}))
}
