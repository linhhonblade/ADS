package main

import (
	"testing"
)

/*
 * Complete the 'pickingNumbers' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY a as parameter.
 */

func pickingNumbers(a []int32) int32 {
	occurences := make(map[int32]int32)
	var max_count, count int32
	for _, v := range a {
		occurences[v]++
	}
	for key, val := range occurences {
		count = val
		if occurences[key+1] >= occurences[key-1] {
			count += occurences[key+1]
		} else {
			count += occurences[key-1]
		}
		if count > max_count {
			max_count = count
		}
	}
	return max_count
}

func TestPickingNumber(t *testing.T) {
	// arr := []int32{1, 1, 2, 2, 4, 4, 5, 5, 5}
	// res := pickingNumbers(arr)
	// if res != 5 {
	// 	t.Errorf("Expected 5, got %v", res)
	// }
	arr := []int32{4, 6, 5, 3, 3, 1}
	res := pickingNumbers(arr)
	if res != 3 {
		t.Errorf("Expected 3, got %v", res)
	}
}
