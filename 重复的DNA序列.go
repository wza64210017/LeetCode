package main

import "fmt"

var m = map[uint8]int{
	'A': 0,
	'C': 1,
	'G': 2,
	'T': 3,
}

func findRepeatedDnaSequences(s string) []string {
	if len(s) <= 10 {
		return nil
	}

	sum := initSum(s[:10])
	sumM := map[int]bool{sum: true}

	set := make(map[string]bool)
	for i := 10; i < len(s); i++ {
		sum = sum<<2 | m[s[i]]
		sum &= 0xfffff

		if sumM[sum] {
			set[s[i-9:i+1]] = true
			continue
		}

		sumM[sum] = true
	}

	res := make([]string, 0, len(set))
	for s := range set {
		res = append(res, s)
	}

	return res
}

func initSum(s string) int {
	sum := 0
	for i := 0; i < len(s); i++ {
		sum = sum<<2 | m[s[i]]
	}

	return sum
}

func main() {
	fmt.Println(findRepeatedDnaSequences("AAAAAAAAAAAA"))
	//fmt.Println(findRepeatedDnaSequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"))
}
