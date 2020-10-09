package main

import (
	"fmt"
)

func getPermutation(n int, k int) string {
	if n == 1 {
		return "1"
	}

	m := []int{1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880}
	arr := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

	k--
	str := ""
	for i := n - 1; i > 0; i-- {
		tmp := k / m[i]
		if k%m[i] == 0 {
			tmp--
			k -= tmp * m[i]
		} else {
			k %= m[i]
		}

		str += arr[tmp]
		arr = append(arr[:tmp], arr[tmp+1:]...)
	}

	str += arr[0]
	return str
}

func main() {
	fmt.Println(getPermutation(9, 2))
}

// 1234

// n! = n * (n-1) * ... * 1

// 先选定第一个: k / (n-1)! = 9 / 6 = 1
//				k % (n-1)! = 9 % 6 = 3

// 1234
// 1243
// 1324
// 1342
// 1423
// 1432

// 2 134
// 2 143

// 2 314
// 2 341

// 2 413
// 2 431

// 再选定第二个: (k % (n-1)!) % (n-2)! = (9 % 6) % 2 = 1
//                                   (11 % 6) % 2 = 3
// 1 34
// 1 43

// 3 14
// 3 41

// 4 13
// 4 31

// 3124
// 3142
// 3214
// 3241
// 3412
// 3421

// 4123
// 4132
// 4213
// 4231
// 4312
// 4321
