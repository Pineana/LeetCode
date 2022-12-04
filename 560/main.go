package main

import "fmt"

func main() {
	fmt.Println(subarraySum([]int{1, 2, 3}, 2))

}

// 穷举
func subarraySum(nums []int, k int) int {
	var ans int
	for i := 0; i < len(nums); i++ {
		sum := nums[i]
		if sum == k {
			ans++
		}
		for j := i + 1; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				ans++
			}
		}
	}
	return ans
}

// 前缀和加上哈希表
func subarraySum2(nums []int, k int) int {
	mp := make(map[int]int)
	mp[0] = 1
	presum, count := 0, 0
	for i := 0; i < len(nums); i++ {
		presum += nums[i]
		if _, ok := mp[presum-k]; ok {
			count++
		}
		mp[presum-k] += 1
	}
	return count
}
