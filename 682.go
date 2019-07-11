package main

import (
	"fmt"
	"strconv"
)

func calPoints(ops []string) int {
	sum := 0
	pos := 0

	scores := make([]int, 0, len(ops))
	for _, op := range ops {

		switch op {
		case "C":
			scores = scores[:pos-1]
			pos--
		case "D":
			score := scores[pos-1]
			scores = append(scores, score*2)
			pos++
		case "+":
			scores = append(scores, scores[pos-1] + scores[pos-2])
			pos++
		default:
			score, _ := strconv.Atoi(op)
			scores = append(scores, score)
			pos++
		}
	}

	for _, score := range scores {
		sum += score
	}
	return sum
}

func main() {
	arr := []string{"5","2","C","D","+"}
	fmt.Println(calPoints(arr))
}
