package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	h := NewHeap(10)
	fmt.Println(h)
	h = append(h, 0)
	upAdjust(h)
	fmt.Println(h)
	heapSort(h)
	fmt.Println(h)
}

type Heap []int

func NewHeap(n int) Heap {
	list := make([]int, 0)
	for i := 0; i < n; i++ {
		list = append(list, rand.Intn(n*10))
	}
	buildHeap(list)
	return list
}

func buildHeap(list []int) Heap {
	parentIndex := (len(list) - 2) / 2
	for i := parentIndex; i >= 0; i-- {
		downAdjust(list, i, len(list)-1)
	}
	return list
}

func downAdjust(list []int, parentIndex int, length int) {
	childIndex := 2*parentIndex + 1
	temp := list[parentIndex]
	for childIndex < length {
		if childIndex+1 < length && list[childIndex+1] < list[childIndex] {
			childIndex = childIndex + 1
		}
		if list[parentIndex] < list[childIndex] {
			break
		}
		list[parentIndex] = list[childIndex]
		parentIndex = childIndex
		childIndex = 2*childIndex + 1
	}
	list[parentIndex] = temp
}

func upAdjust(list []int) {
	childIndex := len(list) - 1
	parentIndex := (childIndex - 1) / 2
	temp := list[childIndex]
	for childIndex != 0 {
		if list[parentIndex] < temp {
			break
		}
		list[childIndex] = list[parentIndex]
		childIndex = parentIndex
		parentIndex = (childIndex - 1) / 2
	}
	list[childIndex] = temp
}

func heapSort(h Heap) {
	for i := len(h) - 1; i > 0; i-- {
		h[0], h[i] = h[i], h[0]
		downAdjust(h, 0, i)
	}
}
