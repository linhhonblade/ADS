package main

import (
	"testing"
)

func groupAnagrams(strs []string) [][]string {
	var res [][]string
	res_map := make(map[[26]int]int)
	for i := 0; i < len(strs); i++ {
		var key [26]int
		for j := 0; j < len(strs[i]); j++ {
			key[strs[i][j]-'a']++
		}
		if ind, ok := res_map[key]; ok {
			res[ind] = append(res[ind], strs[i])
		} else {
			res = append(res, []string{strs[i]})
			res_map[key] = len(res) - 1
		}

	}
	return res
}

func TestGroupAnagram(t *testing.T) {
	var strs []string
	t.Run("Case 1", func(t *testing.T) {
		strs = []string{"eat", "tea", "tan", "ate", "nat", "bat"}
		res := groupAnagrams(strs)
		expected := [][]string{[]string{"bat"}, []string{"nat", "tan"}, []string{"ate", "eat", "tea"}}
		t.Errorf("Expected %v, got %v\n", expected, res)
	})
}
