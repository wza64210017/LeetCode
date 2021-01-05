package 二分法

func search(arr []int, target int) int {
	left, right := 0, len(arr)-1

	// 二分法
	for left < right {
		// 中间值
		mid := (left + right) / 2
		if arr[left] < arr[mid] && arr[left] > arr[mid+1] {
			break
		}


	}
}

func getRight(left int, arr []int) int {
	right := 11
	for left < right {

	}
}

func main() {

}
