package main

import (
	"fmt"
)

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
}

func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(a []int, l, r, index int) int {
	q := partition(a, l, r)
	if q == index {
		return a[q]
	} else if q < index {
		return quickSelect(a, q+1, r, index)
	}
	return quickSelect(a, l, q-1, index)
}

func partition(nums []int, l, r int) int {
	pivotIndex, mark := l, l
	for i := l; i <= r; i++ {
		if nums[i] < nums[pivotIndex] {
			mark++
			nums[i], nums[mark] = nums[mark], nums[i]
		}
	}
	nums[mark], nums[pivotIndex] = nums[pivotIndex], nums[mark]
	return mark
}
