package main

import "testing"

func ShiftedOffsetSearch(shiftArr []int, left, right int) int {
	// base case
	if shiftArr[0] <= shiftArr[len(shiftArr)-1] {
		return -1
	}
	var median int
	for left < right {
		median = (left + right) / 2
		if shiftArr[median] > shiftArr[left] {
			left = median
		} else {
			right = median
		}
	}
	return left
}

func FindItemIndex(shiftArr []int, num, left, right int) int {
	if shiftArr[right] == num {
		return right
	}
	var median int
	for left < right {
		median = (left + right) / 2
		if shiftArr[median] == num {
			return median
		} else if shiftArr[median] > num {
			right = median
		} else if left != median {
			left = median
		} else {
			break
		}
	}
	return -1
}

func ShiftedArrSearch(shiftArr []int, num int) int {
	if len(shiftArr) == 1 {
		if shiftArr[0] == num {
			return 0
		} else {
			return -1
		}
	}
	var res int
	offset := ShiftedOffsetSearch(shiftArr, 0, len(shiftArr)-1)
	if num > shiftArr[len(shiftArr)-1] {
		res = FindItemIndex(shiftArr, num, 0, offset)
	} else {
		res = FindItemIndex(shiftArr, num, offset+1, len(shiftArr)-1)
	}
	return res
}

func TestShiftedArraySearch(t *testing.T) {
	var arr []int
	var res, num int
	t.Run("Case 1", func(t *testing.T) {
		arr = []int{9, 12, 17, 2, 4, 5}
		num = 2
		res = ShiftedArrSearch(arr, num)
		if res != 3 {
			t.Errorf("Expected 3, got %v\n", res)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		arr = []int{14, 16, 17, 19, 20, 3, 5, 9, 10}
		num = 6
		res = ShiftedArrSearch(arr, num)
		if res != -1 {
			t.Errorf("Expected -1, got %v\n", res)
		}
	})
	t.Run("Case 3", func(t *testing.T) {
		arr = []int{2}
		num = 2
		res = ShiftedArrSearch(arr, num)
		if res != 0 {
			t.Errorf("Expected 0, got %v\n", res)
		}
	})
	t.Run("Case 4", func(t *testing.T) {
		arr = []int{1, 2}
		num = 2
		res = ShiftedArrSearch(arr, num)
		if res != 1 {
			t.Errorf("Expected 1, got %v\n", res)
		}
	})
	t.Run("Case 5", func(t *testing.T) {
		arr = []int{0, 1, 2, 3, 4, 5}
		num = 1
		res = ShiftedArrSearch(arr, num)
		if res != 1 {
			t.Errorf("Expected 1, got %v\n", res)
		}
	})
	t.Run("Case 6", func(t *testing.T) {
		arr = []int{1, 2, 3, 4, 5, 0}
		num = 0
		res = ShiftedArrSearch(arr, num)
		if res != 5 {
			t.Errorf("Expected 5, got %v\n", res)
		}
	})
	t.Run("Case 7", func(t *testing.T) {
		arr = []int{9, 12, 17, 2, 4, 5}
		num = 17
		res = ShiftedArrSearch(arr, num)
		if res != 2 {
			t.Errorf("Expected 2, got %v\n", res)
		}
	})
}
