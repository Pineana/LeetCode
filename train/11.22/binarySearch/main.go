package main

import "fmt"

func main() {
	array := buildSortArray(10)

	fmt.Println(binarySearch(array, 9))
}

func buildSortArray(n int) []int {
	array := make([]int, 0)
	for i := 0; i < n; i++ {
		array = append(array, i)
	}
	return array
}

func binarySearch(array []int, item int) int {
	left := 0
	right := len(array) - 1

	for left <= right {
		mid := (left + right) / 2
		if array[mid] > item {
			right = mid - 1
		} else if array[mid] < item {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
