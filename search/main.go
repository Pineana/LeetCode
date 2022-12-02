package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	s1 := buildSortList(10)
	fmt.Println(s1)
	s2 := []int{1, 18, 19, 25, 28, 33, 42, 57, 80, 94}
	fmt.Println(binarySearch(s2, 25))
}

func buildSortList(n int) []int {
	rand.Seed(time.Now().Unix())
	list := make([]int, 0)
	for i := 0; i < n; i++ {
		list = append(list, rand.Intn(100))
	}
	sort.Ints(list)
	return list
}

func binarySearch(sort []int, target int) int {
	l := 0
	r := len(sort) - 1
	for l <= r {
		mid := (l + r) / 2
		if sort[mid] > target {
			r = mid - 1
		} else if sort[mid] < target {
			l = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
