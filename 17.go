package main

import "fmt"

func letterCombinations(digits string) []string {
	if digits == "" {
		return nil
	}

	var Numbers = map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}

	results := make([]string, 1, 16)
	results[0] = ""

	for i := 0; i < len(digits); i++ {
		tmp := make([]string, 0, 16)
		for j := 0; j < len(results); j++ {
			for k := 0; k < len(Numbers[digits[i]]); k++ {
				tmp = append(tmp, results[j]+Numbers[digits[i]][k])
			}
		}

		results = tmp
	}

	return results
}

func main() {
	fmt.Println(letterCombinations("23"))
}
