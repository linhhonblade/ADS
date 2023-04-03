package main

import (
	"testing"
)

func isValidSudoku(board [][]byte) bool {
	row_map, col_map, grid_map := make([]map[byte]byte, 9), make([]map[byte]byte, 9), make([]map[byte]byte, 9)
	var grid_idx int
	for r := 0; r < 9; r++ {
		row_map[r] = make(map[byte]byte)
		col_map[r] = make(map[byte]byte)
		grid_map[r] = make(map[byte]byte)
	}
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] == '.' {
				continue
			}
			if _, ok := row_map[r][board[r][c]]; ok {
				return false
			} else {
				row_map[r][board[r][c]] = 1
			}
			if _, ok := col_map[c][board[r][c]]; ok {
				return false
			} else {
				col_map[c][board[r][c]] = 1
			}
			grid_idx = (r/3)*3 + c/3
			if _, ok := grid_map[grid_idx][board[r][c]]; ok {
				return false
			} else {
				grid_map[grid_idx][board[r][c]] = 1
			}
		}
	}
	return true
}

func TestIsValidSudoku(t *testing.T) {
	var board [][]byte
	t.Run("Case 1", func(t *testing.T) {
		board = [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
			{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
			{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
			{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
			{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
			{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
			{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
			{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
			{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
		expected := true
		res := isValidSudoku(board)
		if res != expected {
			t.Errorf("Expected %v, got %v\n", expected, res)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		board = [][]byte{{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
			{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
			{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
			{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
			{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
			{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
			{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
			{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
			{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
		expected := false
		res := isValidSudoku(board)
		if res != expected {
			t.Errorf("Expected %v, got %v\n", expected, res)
		}

	})
}
