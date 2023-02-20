package main

import "fmt"

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] > pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	i++
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func quickSort(arr []int, low, high int) {
	if low < high {
		arr, q := partition(arr, low, high)
		quickSort(arr, low, q-1)
		quickSort(arr, q+1, high)
	}
}

func main() {
	arr := []int{6, 3, 1, 7, 8, 2, 9, 4, 5}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
