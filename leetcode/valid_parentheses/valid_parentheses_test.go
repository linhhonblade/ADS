package main

import (
	"strings"
	"testing"
)

func isValid(s string) bool {
	var paren []rune
	open_paren := "([{"
	close_paren := ")]}"
	for _, p := range s {
		if strings.Contains(open_paren, string(p)) {
			paren = append(paren, p)
		}
		if strings.Contains(close_paren, string(p)) {
			if len(paren) == 0 {
				return false
			}
			if (p == ')' && paren[len(paren)-1] == '(') ||
				(p == ']' && paren[len(paren)-1] == '[') ||
				(p == '}' && paren[len(paren)-1] == '{') {
				paren = paren[:len(paren)-1]
			} else {
				return false
			}
		}
	}
	if len(paren) == 0 {
		return true
	} else {
		return false
	}
}

func TestIsValidParentheses(t *testing.T) {
	var s string
	var res, expected bool
	t.Run("Case 1", func(t *testing.T) {
		s = "()"
		expected = true
		res = isValid(s)
		if res != expected {
			t.Errorf("Expected %v, got %v\n", expected, res)
		}
	})
	t.Run("Case 1", func(t *testing.T) {
		s = "()[]{}"
		expected = true
		res = isValid(s)
		if res != expected {
			t.Errorf("Expected %v, got %v\n", expected, res)
		}
	})
	t.Run("Case 1", func(t *testing.T) {
		s = "(]"
		expected = false
		res = isValid(s)
		if res != expected {
			t.Errorf("Expected %v, got %v\n", expected, res)
		}
	})
}
