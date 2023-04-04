package main

import (
	"testing"
)

func longestConsecutive(nums []int) int {
	seen := make(map[int]struct{})
	for _, n := range nums {
		seen[n] = struct{}{}
	}
	maxSeqLen := 0
	for n := range seen {
		// If n-1 exists, then n is not the start of the consecutive sequence, ignore and continue
		// Else: n is the start of the consecutive sequence, find the following (if exist) and count
		if _, ok := seen[n-1]; !ok {
			seqLen := 1
			for {
				if _, ok := seen[n+1]; ok {
					seqLen += 1
					n += 1
				} else {
					// this is the end of seq
					maxSeqLen = max(seqLen, maxSeqLen)
					break
				}
			}
		}
	}
	return maxSeqLen
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func TestLongestConsecutive(t *testing.T) {
	var nums []int
	t.Run("Case 1", func(t *testing.T) {
		nums = []int{100, 4, 200, 1, 3, 2}
		expected := 4
		res := longestConsecutive(nums)
		if res != expected {
			t.Errorf("Expected %v, got %v\n", expected, res)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		nums = []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
		expected := 9
		res := longestConsecutive(nums)
		if res != expected {
			t.Errorf("Expected %v, got %v\n", expected, res)
		}
	})
}
