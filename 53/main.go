package main

import "fmt"

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

func maxSubArray(nums []int) int {
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		ans = max(ans, max(nums[i-1]+nums[i], nums[i]))
		nums[i] = max(nums[i-1]+nums[i], nums[i])
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
