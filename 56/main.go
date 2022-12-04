package main

import (
	"fmt"
	"sort"
)

func main() {
	intervals := [][]int{[]int{1, 3}, []int{2, 6}, []int{8, 10}, []int{15, 18}}
	fmt.Println(merge(intervals))
}

func merge(intervals [][]int) [][]int {
	res := make([][]int, 0)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	pre := intervals[0]
	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		if pre[1] < cur[0] {
			res = append(res, pre)
			pre = cur
		} else {
			pre[1] = max(pre[1], cur[1])
		}
	}
	res = append(res, pre)
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
