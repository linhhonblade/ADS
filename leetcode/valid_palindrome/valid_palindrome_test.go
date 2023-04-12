package main

import (
	"log"
	"regexp"
	"strings"
	"testing"
)

func isPalindrome(s string) bool {
	processedStr := sanitizeString(s)
	for i := 0; i < len(processedStr); i++ {
		if i > len(processedStr)-i-1 {
			break
		} else {
			if processedStr[i] != processedStr[len(processedStr)-i-1] {
				return false
			}
		}
	}
	return true
}
func sanitizeString(s string) string {
	// Replace all non-alphanumeric characters with an empty string
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Println(err)
	}
	processedStr := reg.ReplaceAllString(s, "")
	return strings.ToLower(processedStr)
}

func TestValidPalindrome(t *testing.T) {
	var s string
	var expected, res bool
	t.Run("Case 1", func(t *testing.T) {
		s = "A man, a plan, a canal: Panama"
		expected = true
		res = isPalindrome(s)
		if res != expected {
			t.Errorf("Expected %v, got %v\n", expected, res)
		}
	})
	t.Run("Case 2", func(t *testing.T) {
		s = "race a car"
		expected = false
		res = isPalindrome(s)
		if res != expected {
			t.Errorf("Expected %v, got %v\n", expected, res)
		}
	})
	t.Run("Case 3", func(t *testing.T) {
		s = " "
		expected = true
		res = isPalindrome(s)
		if res != expected {
			t.Errorf("Expected %v, got %v\n", expected, res)
		}
	})
}
