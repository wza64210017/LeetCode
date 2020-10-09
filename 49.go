package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string, len(strs)/2)
	for _, str := range strs {
		s := sortStr(str)
		m[s] = append(m[s], str)
	}

	ret := make([][]string, len(m))
	i := 0
	for _, arr := range m {
		ret[i] = make([]string, 0, len(arr))
		for _, v := range arr {
			ret[i] = append(ret[i], v)
		}
		i++
	}

	return ret
}

func sortStr(str string) string {
	b := []byte(str)
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})

	return string(b)
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
