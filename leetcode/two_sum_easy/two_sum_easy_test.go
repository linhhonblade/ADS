package main

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	for i, v := range nums {
		if j, ok := seen[target-v]; ok {
			return []int{i, j}
		}
		seen[v] = i
	}
	return *new([]int)
}
