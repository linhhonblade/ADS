package main

import (
	"testing"
)

/*
consider all possible nodes that can be reached in one step
then expand those nodes to get the range of possibilities for the next step
*/
func jump(nums []int) int {
	res := 0
	l := 0
	r := 0
	for r < len(nums)-1 {
		farthest := 0
		for i := l; i <= r; i++ {
			if i+nums[i] > farthest {
				farthest = i + nums[i]
			}
		}
		l = r + 1
		r = farthest
		res++
	}
	return res
}

func TestJumpGameII(t *testing.T) {
	var res, expected int
	var nums []int
	t.Run("Case 1", func(t *testing.T) {
		nums = []int{2, 3, 1, 1, 4}
		res = jump(nums)
		expected = 2
		if res != expected {
			t.Errorf("Expected %v, got %v\n", expected, res)
		}
	})

	t.Run("Case 2", func(t *testing.T) {
		nums = []int{2, 3, 0, 1, 4}
		res = jump(nums)
		expected = 2
		if res != expected {
			t.Errorf("Expected %v, got %v\n", expected, res)
		}
	})

	t.Run("Case 3", func(t *testing.T) {
		nums = []int{1, 2, 1, 1, 1}
		res = jump(nums)
		expected = 3
		if res != expected {
			t.Errorf("Expected %v, got %v\n", expected, res)
		}
	})
}
