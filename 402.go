package main

import "fmt"

func removeKdigits(num string, k int) string {
	if len(num) == k {
		return "0"
	}


}

func main() {
	str := ""
	k := 0

	fmt.Println(removeKdigits(str, k))
}
