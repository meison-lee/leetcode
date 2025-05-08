package main

import (
	"fmt"
)

// use [9][9] fixed size array to implement fast look up
func isValidSudoku(board [][]byte) bool {
	var row, col, box [9][9]bool

	for i := range board {
		for j := range board[i] {
			if board[i][j] != '.' {
				v := int(board[i][j]-'0') - 1
				if row[i][v] || col[j][v] || box[i/3+(j/3)*3][v] {
					return false
				}

				row[i][v] = true
				col[j][v] = true
				box[i/3+(j/3)*3][v] = true
			}
		}
	}
	return true
}

func main() {
	intput := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	result := isValidSudoku(intput)
	fmt.Println(result)
}
