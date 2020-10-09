package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	ls1, ls2 := len(s1), len(s2)
	if ls1 > ls2 {
		return false
	}

	arrS1, arrS2 := [26]int{}, [26]int{}
	for i := range s1 {
		arrS1[s1[i]-'a']++
		arrS2[s2[i]-'a']++
	}

	for i := 0; i < ls2-ls1; i++ {
		if arrS1 == arrS2 {
			return true
		}

		arrS2[s2[i]-'a']--
		arrS2[s2[ls1+i]-'a']++
	}

	return arrS1 == arrS2
}

func main() {
	fmt.Println(checkInclusion("adc", "dcda"))
}

// a d c d e
// a d c d d a
