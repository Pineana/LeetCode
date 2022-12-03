package main

import "sort"

func main() {

}

func threeSum(nums []int) [][]int {
	ans := make([][]int, 0)
	if nums == nil || len(nums) < 3 {
		return ans
	}
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			if nums[left]+nums[right]+nums[i] == 0 {
				ans = append(ans, []int{nums[left], nums[right], nums[i]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if nums[left]+nums[right]+nums[i] < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return ans
}
