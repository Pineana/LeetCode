package main

import "fmt"

func main() {
	array := buildSortArray(10)

	fmt.Println(binarySearch(array, 3))
}

func buildSortArray(n int) []int {
	array := make([]int, 0)
	for i := 0; i < n; i++ {
		array = append(array, i)
	}
	return array
}

func binarySearch(array []int, item int) int {
	L := 0
	R := len(array) - 1
	for L <= R {
		mid := (L + R) / 2
		if array[mid] > item {
			R = mid - 1
		} else if array[mid] < item {
			L = mid + 1
		} else {
			return mid
		}
	}
	return -1

}
