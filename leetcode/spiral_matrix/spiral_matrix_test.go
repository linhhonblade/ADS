package main

import (
	slices "golang.org/x/exp/slices"
	"testing"
)

func trimMatrix(matrix [][]int, mode int8) ([]int, [][]int) {
	/*
		0: top
		1: right
		2: bottom
		3: left
	*/
	res := []int{}
	lastCol := len(matrix[0]) - 1
	lastRow := len(matrix) - 1
	if mode == 0 {
		res = matrix[0]
		if lastRow == 0 {
			matrix = matrix[:0]
		} else {
			matrix = matrix[1:]
		}
	}
	if mode == 1 {
		for i := 0; i <= lastRow; i++ {
			res = append(res, matrix[i][lastCol])
			matrix[i] = matrix[i][:lastCol]
		}
	}
	if mode == 2 {
		for i := lastCol; i >= 0; i-- {
			res = append(res, matrix[lastRow][i])
		}
		matrix = matrix[:lastRow]
	}
	if mode == 3 {
		for i := lastRow; i >= 0; i-- {
			res = append(res, matrix[i][0])
			if lastCol == 0 {
				matrix[i] = matrix[i][:0]
			} else {
				matrix[i] = matrix[i][1:]
			}
		}
	}
	return res, matrix
}

func spiralOrder(matrix [][]int) []int {
	var res []int
	var count int8
	var side []int
	for len(matrix) > 0 && len(matrix[0]) > 0 {
		side, matrix = trimMatrix(matrix, count%4)
		res = append(res, side...)
		count++
	}
	return res
}

func TestSpiralMatrix(t *testing.T) {
	// a := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	// expected := []int{1, 2, 3, 6, 9, 8, 7, 4, 5}
	// a := [][]int{{7}, {9}, {6}}
	// expected := []int{7, 9, 6}
	a := [][]int{{0, 1, 2, 3, 4}, {5, 6, 7, 8, 9}, {10, 11, 12, 13, 14}}
	expected := []int{0, 1, 2, 3, 4, 9, 14, 13, 12, 11, 10, 5, 6, 7, 8}
	t.Run("Spiral Matrix 1", func(t *testing.T) {
		res := spiralOrder(a)
		if !slices.Equal(res, expected) {
			t.Errorf("Expected %v, got %v\n", expected, res)
		}
	})
}
