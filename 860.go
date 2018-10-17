package main

import "fmt"

func lemonadeChange(bills []int) bool {

	five := 0
	ten := 0
	for _, bill := range bills {
		switch bill {
		case 5:
			five++
		case 10:
			if five == 0 {
				return false
			}

			five--
			ten++
		case 20:
			if ten == 0 {
				if five < 3 {
					return false
				}

				five -= 3
			} else {
				if five == 0 {
					return false
				}

				ten--
				five--
			}
		}
	}

	return true
}

//柠檬水找零
//[5,5,5,10,20] true
func main() {
	fmt.Println(lemonadeChange([]int{5,5,10,10,20}))
}
