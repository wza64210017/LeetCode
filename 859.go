package main

import "fmt"

func buddyStrings(A string, B string) bool {
	lengthA := len(A)
	lengthB := len(B)

	if lengthA != lengthB || lengthA == 0 {
		return false
	}

	if A == B {
		tmp := map[string]int{}
		for _, v := range A {
			if tmp[string(v)] != 0 {
				return true
			}

			tmp[string(v)] = 1
		}

		return false
	}

	s := -1
	for k := range A {
		if A[k] == B[k] {
			continue
		}

		if s == -1 {
			s = k
		} else {
			if A[s] != B[k] || A[k] != B[s] {
				return false
			}
		}
	}

	return true
}

//亲密字符串
func main() {
	fmt.Println(buddyStrings("", "aa"))
}
