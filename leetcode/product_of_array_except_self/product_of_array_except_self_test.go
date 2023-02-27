package main

import (
	slices "golang.org/x/exp/slices"
	"testing"
)

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	right := 1
	res[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		res[i] = res[i-1] * nums[i]
	}
	for i := len(res) - 1; i > 0; i-- {
		res[i] = res[i-1] * right
		right = right * nums[i]
	}
	res[0] = right
	return res
}

func TestProductExceptSelf(t *testing.T) {
	var nums, expected, res []int
	t.Run("Case 1", func(t *testing.T) {
		nums = []int{1, 2, 3, 4}
		expected = []int{24, 12, 8, 6}
		res = productExceptSelf(nums)
		if !slices.Equal(res, expected) {
			t.Errorf("Expected %v, got %v\n", expected, res)
		}
	})
}
