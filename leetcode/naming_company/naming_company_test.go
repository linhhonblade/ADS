package main

import (
	"testing"
)

func distinctNames(ideas []string) int64 {
	var count int64
	group := make(map[string]map[string]bool)
	for _, s := range ideas {
		if group[string(s[1:])] == nil {
			group[string(s[1:])] = make(map[string]bool)
		}
		if _, ok := group[string(s[1:])][string(s[0])]; !ok {
			group[string(s[1:])][string(s[0])] = true
		}
	}
	for i, s := range ideas {
		for j := i + 1; j < len(ideas); j++ {
			_, ok1 := group[string(s[1:])][string(ideas[j][0])]
			_, ok2 := group[string(ideas[j][1:])][string(s[0])]
			if ok1 || ok2 {
				continue
			}
			count += 2
		}
	}
	return count
}

func TestDistinctName(t *testing.T) {
	var ideas []string
	var expected, res int64
	t.Run("Case 1", func(t *testing.T) {
		ideas = []string{"coffee", "donuts", "time", "toffee"}
		res = distinctNames(ideas)
		expected = 6
		if res != expected {
			t.Errorf("Expected %v, got %v\n", expected, res)
		}
	})

	t.Run("Case 2", func(t *testing.T) {
		ideas = []string{"bzklqtbdr", "kaqvdlp", "r", "dk"}
		res = distinctNames(ideas)
		expected = 12
		if res != expected {
			t.Errorf("Expected %v, got %v\n", expected, res)
		}
	})
}
