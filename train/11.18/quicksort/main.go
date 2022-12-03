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
	pivotIndex := paratationDouble(array, L, R)
	quickSort(array, L, pivotIndex-1)
	quickSort(array, pivotIndex+1, R)
}

func paratation(array []int, L, R int) int {
	pivotIndex := L
	mark := L
	for i := L; i <= R; i++ {
		if array[i] < array[pivotIndex] {
			mark++
			array[mark], array[i] = array[i], array[mark]
		}
	}
	array[pivotIndex], array[mark] = array[mark], array[pivotIndex]
	return mark
}

func paratationDouble(array []int, L, R int) int {
	pivotIndex := L
	for L < R {
		for array[R] > array[pivotIndex] && L < R {
			R--
		}
		for array[L] < array[pivotIndex] && L < R {
			L++
		}
		array[L], array[R] = array[R], array[L]
	}
	array[pivotIndex], array[L] = array[L], array[pivotIndex]
	return L

}
