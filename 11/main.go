package main

func main() {

}

func maxArea(height []int) int {
	left, right, ans := 0, len(height)-1, 0
	for left < right {
		area := min(height[left], height[right]) * (right - left)
		ans = max(area, ans)
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
	} else {
		return y
	}
}

func min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}
