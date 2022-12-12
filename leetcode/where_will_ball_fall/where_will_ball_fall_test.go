package main

import (
	slices "golang.org/x/exp/slices"
	"testing"
)

func getNextCell([]int prev, []int current, k) {
	if current[0] != prev[0] {
		// downward
		return []int{}
	}
}

func findBall(grid [][]int) []int {
	var res, history []int
	prev_row := -1
	current_row := 0
	max_row := len(grid) - 1
	max_col := len(grid[0]) - 1
	for current_row <= max_row {
		if current_row != prev_row {
			// here only shift horizontally
			prev_row = current_row
			history = res
			for col := 0; col <= max_col; col++ {
				if res[col] != -1 {
					res[col] += grid[current_row][col]
				}
				if res[col] > max_col {
					res[col] = -1
				}

			}
		}
	}
	return res

}

func TestFindBall(t *testing.T) {
	grid := [][]int{{1, 1, 1, -1, -1}, {1, 1, 1, -1, -1}, {-1, -1, -1, 1, 1}, {1, 1, 1, 1, -1}, {-1, -1, -1, -1, -1}}
	expected := []int{1, -1, -1, -1, -1}
	res := findBall(grid)
	if !slices.Equal(res, expected) {
		t.Errorf("Expected %v, got %v\n", expected, res)
	}
}
