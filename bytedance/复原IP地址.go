package main

import (
	"fmt"
	"strconv"
	"strings"
)

var result []string

func restoreIpAddresses(s string) []string {
	if len(s) < 4 || len(s) > 12 {
		return nil
	}

	result = make([]string, 0)
	backTrack(s, []string{}, 1)

	return result
}

func backTrack(s string, track []string, times int) {
	if times == 5 && len(s) == 0 {
		str := strings.Join(track, ".")
		result = append(result, str)
		return
	}

	// 每段长度 1 ～ 3
	for i := 1; i <= 3; i++ {
		if i <= len(s) {
			if valid(s[:i]) {
				track = append(track, s[:i])
				backTrack(s[i:], track, times+1)
				track = track[:len(track)-1]
			}

			if s[:i] == "0" {
				break
			}
		}
	}
}

func valid(ip string) bool {
	intIP, _ := strconv.Atoi(ip)
	return intIP >= 0 && intIP <= 255
}

func main() {
	fmt.Println(restoreIpAddresses("25525511135"))
}
