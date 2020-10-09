package main

import "fmt"

func getLeastNumbers(arr []int, k int) []int {
	quickSort(arr, 0, len(arr)-1)
	return arr[0:k]
}

func quickSort(arr []int, left, right int) {
	if left >= right {
		return
	}

	i, j := left, right
	middle := arr[(left+right)/2]

	for i <= j {
		for arr[i] < middle {
			i++
		}

		for arr[j] > middle {
			j--
		}

		if i <= j {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
	}

	for i <= j && arr[i] < middle {
		i++
	}

	for i <= j && arr[j] > middle {
		j--
	}

	if left < j {
		quickSort(arr, left, j)
	}

	if right > i {
		quickSort(arr, i, right)
	}
}

func main() {
	fmt.Println(getLeastNumbers([]int{3, 2, 1}, 2))
}
