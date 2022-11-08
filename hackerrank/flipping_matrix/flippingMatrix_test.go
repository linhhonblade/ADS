package main

import (
	"testing"
)

func flippingMatrix(arr [][]int) int {
	var max, res int
	size := len(arr[0])
	for i := 0; i < size/2; i++ {
		for j := 0; j < size/2; j++ {
			max = arr[i][j]
			if arr[i][size-1-j] > max {
				max = arr[i][size-1-j]
			}
			if arr[size-1-i][j] > max {
				max = arr[size-1-i][j]
			}
			if arr[size-1-i][size-1-j] > max {
				max = arr[size-1-i][size-1-j]
			}
			res += max
		}
	}
	return res
}

func TestFlippingMatrix(t *testing.T) {
	arr := [][]int{
		{112, 42, 83, 119},
		{56, 125, 56, 49},
		{15, 78, 101, 43},
		{62, 98, 114, 108},
	}
	t.Run("FlippingMatrix", func(t *testing.T) {
		res := flippingMatrix(arr)
		if res != 414 {
			t.Errorf("Expected 414, got %d\n", res)
		}
	})
}
