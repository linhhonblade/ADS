package main

import "testing"

func candy(ratings []int) int {
	candies := make([]int, len(ratings))
	for i, _ := range candies {
		candies[i] = 1
	}
	for i := 1; i < len(candies); i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}
	for i := len(candies) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			if candies[i+1]+1 > candies[i] {
				candies[i] = candies[i+1] + 1
			}
		}
	}
	sum := 0
	for _, v := range candies {
		sum += v
	}
	return sum
}

func TestCandy(t *testing.T) {
	var ratings []int
	t.Run("Case 1", func(t *testing.T) {
		ratings = []int{1, 0, 2}
		res := candy(ratings)
		expected := 5
		if res != expected {
			t.Errorf("Expected %d, got %d\n", expected, res)
		}
	})

	t.Run("Case 2", func(t *testing.T) {
		ratings = []int{1, 2, 2}
		res := candy(ratings)
		expected := 4
		if res != expected {
			t.Errorf("Expected %d, got %d\n", expected, res)
		}
	})
	t.Run("Case 3", func(t *testing.T) {
		ratings = []int{1, 2, 87, 87, 87, 2, 1}
		res := candy(ratings)
		expected := 13
		if res != expected {
			t.Errorf("Expected %d, got %d\n", expected, res)
		}
	})
}
