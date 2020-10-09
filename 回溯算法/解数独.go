package main

import "fmt"

var success bool
var row, col, box [9]int

func solveSudoku(board [][]byte) {
	row, col, box, success = [9]int{}, [9]int{}, [9]int{}, false // 状态重置
	initArr(board)

	backtrace(board, row, col, box, 0)
	for _, arr := range board {
		for _, v := range arr {
			fmt.Print(v-'0', " ")
		}
		fmt.Println()
	}
}

func initArr(board [][]byte) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] != '.' {
				tmp := 0x1 << (board[i][j] - '0')
				idx := i/3*3 + j/3

				row[i] |= tmp
				col[j] |= tmp
				box[idx] |= tmp
			}
		}
	}
}

func backtrace(board [][]byte, row, col, box [9]int, times int) {
	// 超过81次
	if times >= 81 {
		success = true
		return
	}

	r := times / 9 // 当前行数
	c := times % 9 // 当前列数
	if board[r][c] != '.' {
		backtrace(board, row, col, box, times+1)
		return
	}

	idx := r/3*3 + c/3 // 当前3x3数

	var i byte
	for i = 1; i <= 9; i++ {
		tmp := 0x1 << i
		// 判断该数在行、列、3x3区间是否存在
		if row[r]&tmp > 0 || col[c]&tmp > 0 || box[idx]&tmp > 0 {
			continue
		}

		// 填充数据
		row[r] |= tmp
		col[c] |= tmp
		box[idx] |= tmp
		board[r][c] = i + '0'

		backtrace(board, row, col, box, times+1)
		if success {
			break
		}

		// 移除
		row[r] ^= tmp
		col[c] ^= tmp
		box[idx] ^= tmp
		board[r][c] = '.'
	}
}

//func solveSudoku(board [][]byte) {
//	rowUsed := make([]map[byte]bool, 9)
//	colUsed := make([]map[byte]bool, 9)
//	rangeUsed := make([]map[byte]bool, 9)
//
//	initMap(&rowUsed)
//	initMap(&colUsed)
//	initMap(&rangeUsed)
//
//	for i := 0; i < 9; i++ {
//		for j := 0; j < 9; j++ {
//			num := board[i][j]
//			if num != '.' {
//				rowUsed[i][num] = true
//				colUsed[j][num] = true
//				box := calRange(i, j)
//				rangeUsed[ran][num] = true
//			}
//		}
//	}
//
//	backtrace(board, 0, 0, &rowUsed, &colUsed, &rangeUsed)
//	for i := 0; i < len(board); i++ {
//		for j := 0; j < len(board[0]); j++ {
//			fmt.Print(board[i][j]-'0', " ")
//		}
//		fmt.Println()
//	}
//}
//
//func initMap(m *[]map[byte]bool) {
//	for k := range *m {
//		(*m)[k] = make(map[byte]bool, 9)
//	}
//}
//
//func calRange(row, col int) int {
//	return (row/3)*3 + col/3
//}
//
//var success bool
//
//func backtrace(board [][]byte, row, col int, rowUsed, colUsed, rangeUsed *[]map[byte]bool) {
//	if row > 8 || col > 8 {
//		success = true
//		return
//	}
//
//	num := (board)[row][col]
//	box := calRange(row, col)
//	if num != '.' {
//		(*rowUsed)[row][num] = true
//		(*colUsed)[col][num] = true
//		(*rangeUsed)[ran][num] = true
//		if col == 8 {
//			backtrace(board, row+1, 0, rowUsed, colUsed, rangeUsed)
//		} else {
//			backtrace(board, row, col+1, rowUsed, colUsed, rangeUsed)
//		}
//
//		return
//	}
//
//	for num = 1 + '0'; num <= 9+'0'; num++ {
//		// 当前数字在该行被使用
//		if (*rowUsed)[row][num] == true {
//			continue
//		}
//
//		// 当前数据在该列被使用
//		if (*colUsed)[col][num] == true {
//			continue
//		}
//
//		// 当前数字在该3x3宫格被使用
//		if (*rangeUsed)[ran][num] == true {
//			continue
//		}
//
//		(*rowUsed)[row][num] = true
//		(*colUsed)[col][num] = true
//		(*rangeUsed)[ran][num] = true
//
//		(board)[row][col] = num
//
//		if col == 8 {
//			backtrace(board, row+1, 0, rowUsed, colUsed, rangeUsed)
//		} else {
//			backtrace(board, row, col+1, rowUsed, colUsed, rangeUsed)
//		}
//
//		if success == true {
//			break
//		}
//
//		(board)[row][col] = '.'
//		(*rowUsed)[row][num] = false
//		(*colUsed)[col][num] = false
//		(*rangeUsed)[ran][num] = false
//	}
//}

func main() {
	board := [][]byte{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}

	board = [][]byte{
		{'.', '.', '9', '7', '4', '8', '.', '.', '.'},
		{'7', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '2', '.', '1', '.', '9', '.', '.', '.'},
		{'.', '.', '7', '.', '.', '.', '2', '4', '.'},
		{'.', '6', '4', '.', '1', '.', '5', '9', '.'},
		{'.', '9', '8', '.', '.', '.', '3', '.', '.'},
		{'.', '.', '.', '8', '.', '3', '.', '2', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '6'},
		{'.', '.', '.', '2', '7', '5', '9', '.', '.'},
	}

	solveSudoku(board)
}
