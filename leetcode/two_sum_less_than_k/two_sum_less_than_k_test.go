package main

import (
	"sort"
	"testing"
)

func TwoSumLessThanK(A []int, K int) int {
	sort.Slice(A, func(i, j int) bool { return A[i] < A[j] })
	i := 0
	j := len(A) - 1
	res := -1
	for i < j {
		if A[j]+A[i] > K {
			j--
		} else {
			if A[i]+A[j] > res {
				res = A[i] + A[j]
			}
			i++
		}
	}
	return res
}

func TestTwoSumLessThanK(t *testing.T) {
	var A []int
	var K, expected int
	t.Run("Case 1", func(t *testing.T) {
		A = []int{34, 23, 1, 24, 75, 33, 54, 8}
		K = 60
		expected = 58
		res := TwoSumLessThanK(A, K)
		if res != expected {
			t.Errorf("Expected %v, got %v\n", expected, res)
		}
	})
}
