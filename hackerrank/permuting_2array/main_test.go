package main

import (
	"sort"
	"testing"
)

/*
 * Complete the 'twoArrays' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. INTEGER k
 *  2. INTEGER_ARRAY A
 *  3. INTEGER_ARRAY B
 */

func twoArrays(k int32, A []int32, B []int32) string {
	// Write your code here
	sort.Slice(A, func(i, j int) bool { return A[i] < A[j] })
	sort.Slice(B, func(i, j int) bool { return B[i] > B[j] })
	for i := 0; i < len(A); i++ {
		if A[i]+B[i] < k {
			return "NO"
		}
	}
	return "YES"
}

func TestTwoArrays(t *testing.T) {
	A := []int32{24, 43, 3, 27, 18, 4, 28, 16, 14, 32, 8, 42, 29, 39, 36, 16, 38, 12, 9, 42, 28, 29, 17, 40, 31, 3, 18, 37, 33, 33, 12, 41, 1, 44, 43, 9, 20, 7}
	B := []int32{49, 38, 6, 48, 11, 44, 38, 24, 9, 10, 31, 23, 25, 11, 2, 32, 11, 37, 30, 14, 30, 3, 20, 16, 30, 12, 33, 43, 31, 23, 16, 51, 43, 47, 0, 21, 9, 21}
	res := twoArrays(44, A, B)
	if res != "YES" {
		t.Errorf("Meow")
	}
}
