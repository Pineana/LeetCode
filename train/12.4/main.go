package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea(height []int) int {
	left, right, ans := 0, len(height)-1, 0
	for left < right {
		h := min(height[left], height[right])
		ans = max(ans, h*(right-left))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
