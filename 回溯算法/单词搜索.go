package main

import "fmt"

func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}

	x, y := len(board), len(board[0])
	used := make([][]bool, x)
	for i := 0; i < x; i++ {
		used[i] = make([]bool, y)
	}

	for i := 0; i < x; i ++ {
		for j := 0; j < y; j++ {
			if board[i][j] != word[0] {
				continue
			}

			used[i][j] = true
			if backtrace(board, used, i, j, word[1:]) == true {
				return true
			}

			used[i][j] = false
		}
	}

	return false
}

var scopes = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func backtrace(board [][]byte, used [][]bool, x, y int, word string) bool {
	if word == "" {
		return true
	}

	var i, j int
	for _, scope := range scopes {
		i, j = x+scope[0], y+scope[1]
		// 行未越界
		validI := i >= 0 && i < len(board)
		// 列未越界
		validJ := j >= 0 && j < len(board[0])

		if validI && validJ && !used[i][j] && board[i][j] == word[0] {
			used[i][j] = true
			if backtrace(board, used, i, j, word[1:]) {
				return true
			}

			used[i][j] = false
		}
	}

	return false
}

func main() {
	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}

	fmt.Println(exist(board, "ABCCED"))
	fmt.Println(exist(board, "SEE"))
	fmt.Println(exist(board, "ABCB"))
	fmt.Println(exist(board, "BFDASA"))
}
