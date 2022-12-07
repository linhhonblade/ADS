package main

import "testing"

func isHappy(n int) bool {
	seen := map[int]bool{}
	for n != 1 {
		val := 0
		for n != 0 {
			v := n % 10
			val += v * v
			n /= 10
		}
		if _, ok := seen[val]; ok {
			return false
		}
		seen[val] = true
		n = val
	}
	return true
}

func TestHappyNumber(t *testing.T) {
	n := 19
	t.Run("Happy Number", func(t *testing.T) {
		res := isHappy(n)
		if !res {
			t.Errorf("Expected true (%d is happy number), got false\n", n)
		}
	})
}
