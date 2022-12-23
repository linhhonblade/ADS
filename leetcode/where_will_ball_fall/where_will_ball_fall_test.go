package main

import (
	slices "golang.org/x/exp/slices"
	"testing"
)

func findBall(grid [][]int) []int {
	var res, history []int
	prev_row := -1
	current_row := 0
	max_row := len(grid) - 1
	max_col := len(grid[0]) - 1
	for i := 0; i <= max_col; i++ {
		res = append(res, i)
		history = append(history, i)
	}
	for current_row <= max_row {
		if current_row != prev_row {
			// here only shift horizontally
			prev_row = current_row
			for col := 0; col <= max_col; col++ {
				if res[col] != -1 {
					history[col] = res[col]
					res[col] += grid[current_row][res[col]]
				}
				if res[col] > max_col {
					res[col] = -1
				}
			}
		} else {
			// here only ship downward
			for col := 0; col <= max_col; col++ {
				if res[col] != -1 {
					if res[col] > history[col] && grid[current_row][res[col]] == -1 {
						res[col] = -1
					} else if res[col] < history[col] && grid[current_row][res[col]] == 1 {
						res[col] = -1
					}
				}

			}
			current_row += 1
		}
	}
	return res

}

func setCellValue(grid [][]int, i, j int8) {
	if grid[i][j] < 2 {
		jNext := j + int8(grid[i][j])
		if jNext < 0 || jNext == int8(len(grid[0])) {
			grid[i][j] = 2
		} else if grid[i][j]+grid[i][jNext] == 0 {
			grid[i][j] = 2
			grid[i][jNext] = 2
		} else {
			setCellValue(grid, i+1, jNext)
			grid[i][j] = grid[i+1][jNext]
		}
	}
}

func findBall11ms(grid [][]int) []int {
	n := len(grid[0])
	lastRow := make([]int, n)
	for j := 0; j < n; j++ {
		lastRow[j] = j + 3
	}
	grid = append(grid, lastRow)
	for j, jEnd := int8(0), int8(n); j < jEnd; j++ {
		setCellValue(grid, 0, j)
	}
	for j := 0; j < n; j++ {
		grid[0][j] -= 3
	}
	return grid[0]
}

func TestFindBall(t *testing.T) {
	grid := [][]int{{1, 1, 1, -1, -1}, {1, 1, 1, -1, -1}, {-1, -1, -1, 1, 1}, {1, 1, 1, 1, -1}, {-1, -1, -1, -1, -1}}
	expected := []int{1, -1, -1, -1, -1}
	// grid := [][]int{{-1}}
	// expected := []int{-1}
	// grid := [][]int{{1, 1, 1, 1, 1, 1}, {-1, -1, -1, -1, -1, -1}, {1, 1, 1, 1, 1, 1}, {-1, -1, -1, -1, -1, -1}}
	// expected := []int{0, 1, 2, 3, 4, -1}
	// grid := [][]int{{1, -1, -1, 1, -1, 1, 1, 1, 1, 1, -1, 1, 1, 1, 1, 1, 1, -1, -1, -1, -1, -1, -1, 1, -1, 1, -1, 1, -1, -1, -1, -1, 1, -1, 1, 1, -1, -1, -1, -1, -1, 1}, {-1, 1, 1, 1, -1, -1, -1, -1, 1, 1, 1, -1, -1, -1, 1, -1, -1, 1, 1, 1, 1, 1, 1, -1, 1, -1, -1, -1, -1, -1, 1, -1, 1, -1, -1, -1, -1, 1, 1, -1, 1, 1}, {1, -1, -1, -1, -1, 1, -1, 1, 1, 1, 1, 1, 1, 1, -1, 1, -1, -1, -1, 1, -1, -1, 1, -1, 1, -1, 1, -1, -1, 1, -1, 1, -1, 1, 1, -1, -1, 1, 1, -1, 1, -1}}
	// expected := []int{-1, -1, 1, -1, -1, -1, -1, 10, 11, -1, -1, 12, 13, -1, -1, -1, -1, -1, 17, -1, -1, 20, -1, -1, -1, -1, -1, -1, -1, -1, 27, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}
	res := findBall(grid)
	if !slices.Equal(res, expected) {
		t.Errorf("Expected %v, got %v\n", expected, res)
	}
}
