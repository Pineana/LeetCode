package main

import "fmt"

func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

func trap(height []int) int {
	z1 := make([]int, len(height))
	z2 := make([]int, len(height))
	ans := 0
	z1[0], z2[len(height)-1] = height[0], height[len(height)-1]
	for i := 1; i < len(height); i++ {
		z1[i] = max(z1[i-1], height[i])
	}

	for i := len(height) - 2; i >= 0; i-- {
		z2[i] = max(z2[i+1], height[i])
	}
	fmt.Println(z1)
	fmt.Println(z2)
	fmt.Println(height)
	for i := 0; i < len(height); i++ {
		ans += min(z1[i], z2[i]) - height[i]
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
