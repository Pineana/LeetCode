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

func quickSort(array []int, L, R int) {
	if L >= R {
		return
	}
	pivotIndex := partitionDouble(array, L, R)
	quickSort(array, L, pivotIndex-1)
	quickSort(array, pivotIndex+1, R)
}

func partition(array []int, L, R int) int {
	pivotIndex := L
	mark := L
	for i := L; i <= R; i++ {
		if array[pivotIndex] < array[i] {
			mark++
			array[mark], array[i] = array[i], array[mark]
		}
	}
	array[pivotIndex], array[mark] = array[mark], array[pivotIndex]
	return mark
}

func partitionDouble(array []int, L, R int) int {
	pivotIndex := L
	for L < R {
		for L < R && array[R] < array[pivotIndex] {
			R--
		}
		for L < R && array[L] >= array[pivotIndex] {
			L++
		}
		array[L], array[R] = array[R], array[L]
	}
	array[L], array[pivotIndex] = array[pivotIndex], array[L]
	return L
}
