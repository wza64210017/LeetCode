package main

import (
	"fmt"
	"strings"
)

var m map[string]bool

func permutation(s string) []string {
	m = make(map[string]bool, 0)
	track := make([]string, 0)
	backtrace(s, track)

	result := make([]string, 0, len(m))
	for k := range m {
		result = append(result, k)
	}

	return result
}

func backtrace(s string, track []string) {
	// 已选解集合满足条件
	if len(s) == 0 {
		sub := strings.Join(track, "")
		m[sub] = true
		return
	}

	// 遍历每个阶段的可选解集合
	for i := range s {
		// 选择此阶段其中一个解,将其加入到已选解集合中
		track = append(track, string(s[i]))
		// 进入下一个阶段
		backtrace(s[0:i]+s[i+1:], track)
		// 「回溯」换个解再遍历
		track = track[:len(track)-1]
	}
}

func main() {
	fmt.Println(permutation("aab"))
}
