package main

import "testing"

func lengthOfLongestSubstring(s string) int {
	seen := make(map[rune]struct{})
	start := 0
	maxLen := 0
	for end, c := range s {
		if _, ok := seen[c]; !ok {
			seen[c] = struct{}{}
			end += 1
			if end-start > maxLen {
				maxLen = end - start
			}
		} else {
			for i, ss := range s[start:end] {
				if ss != c {
					delete(seen, ss)
					continue
				} else if ss == c {
					start = start + i + 1
					break
				}
			}

		}
	}
	return maxLen
}

func TestLenghtOfLongestSubstring(t *testing.T) {
	var s string
	t.Run("Case 1", func(t *testing.T) {
		s = "abcabcbb"
		res := lengthOfLongestSubstring(s)
		expected := 3
		if res != expected {
			t.Errorf("Expected %v, got %v\n", expected, res)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		s = "bbbbb"
		res := lengthOfLongestSubstring(s)
		expected := 1
		if res != expected {
			t.Errorf("Expected %v, got %v\n", expected, res)
		}
	})
	t.Run("Case 3", func(t *testing.T) {
		s = "pwwkew"
		res := lengthOfLongestSubstring(s)
		expected := 3
		if res != expected {
			t.Errorf("Expected %v, got %v\n", expected, res)
		}
	})
	t.Run("Case 4", func(t *testing.T) {
		s = "tmmzuxt"
		res := lengthOfLongestSubstring(s)
		expected := 5
		if res != expected {
			t.Errorf("Expected %v, got %v\n", expected, res)
		}
	})
}
