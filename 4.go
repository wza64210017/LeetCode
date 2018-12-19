package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	l1 := len(nums1)
	l2 := len(nums2)
	i, j := 0, 0

	length := l1 + l2
	newArr := make([]int, (length/2)+1)

	for {
		if i+j > length/2 {
			break
		}

		if i >= l1 {
			newArr[i+j] = nums2[j]
			j++
		} else if j >= l2 {
			newArr[i+j] = nums1[i]
			i++
		} else if nums1[i] < nums2[j] {
			newArr[i+j] = nums1[i]
			i++
		} else {
			newArr[i+j] = nums2[j]
			j++
		}
	}

	if length%2 == 0 {
		return float64(newArr[length/2]+newArr[length/2-1]) / 2
	}

	return float64(newArr[length/2])
}

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}

	fmt.Print(findMedianSortedArrays(nums1, nums2))
}
