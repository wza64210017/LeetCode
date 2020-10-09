package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	tmp := make([]int, m+n)
	copy(tmp, nums1)

	j, k := 0, 0
	for i := 0; i < m+n; i++ {
		if k >= n {
			nums1[i] = tmp[j]
			j++
			continue
		}

		if j >= m {
			nums1[i] = nums2[k]
			k++
			continue
		}

		if tmp[j] < nums2[k] {
			nums1[i] = tmp[j]
			j++
		} else {
			nums1[i] = nums2[k]
			k++
		}
	}
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}

	merge(nums1, 3, nums2, 3)
	fmt.Println(nums1)
}
