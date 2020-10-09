package 回溯算法

import "fmt"

func permutation(S string) []string {
	ret := make([]string, 0)
	use := make(map[int]bool)
	backtrace(S, "", use, &ret)

	return ret
}

func backtrace(s, sub string, use map[int]bool, ret *[]string) {
	if len(s) == len(sub) {
		*ret = append(*ret, sub)
		return
	}

	for i := 0; i < len(s); i++ {
		if use[i] == true {
			continue
		}

		use[i] = true
		backtrace(s, sub+string(s[i]), use, ret)
		delete(use, i)
	}
}

func main() {
	fmt.Println(permutation("abc"))
}
