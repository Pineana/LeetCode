package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	array := buildArray(10)
	fmt.Println(array)
	quickSort(array, 0, len(array)-1)
	fmt.Println(array)
}

func buildArray(n int) []int {
	array := make([]int, 0)
	for i := 0; i < n; i++ {
		array = append(array, rand.Intn(100))
	}
	return array
}

func quickSort(s []int, left, right int) {
	if left >= right {
		return
	}
	pivotIndex := paritation(s, left, right)
	quickSort(s, left, pivotIndex-1)
	quickSort(s, pivotIndex+1, right)

}

func paritation(s []int, left, right int) int {
	pivotIndex := left
	mark := left
	for i := left; i <= right; i++ {
		if s[i] > s[pivotIndex] {
			mark++
			s[mark], s[i] = s[i], s[mark]
		}
	}
	s[pivotIndex], s[mark] = s[mark], s[pivotIndex]
	return mark
}
