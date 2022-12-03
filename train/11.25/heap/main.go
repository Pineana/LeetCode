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
	// h = append(h, 0)
	// upAdjust(h)
	// fmt.Println(h)
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

func buildHeap(h Heap) {
	parentIndex := (len(h) - 2) / 2
	for i := parentIndex; i >= 0; i-- {
		downAssert(h, i, len(h))
	}
}

func heapSort(h Heap) {
	for i := len(h)-1; i > 0; i-- {
		h[0], h[i] = h[i], h[0]
		downAssert(h, 0, i)
	}
}

func downAssert(h Heap, parentIndex int, length int) {
	childIndex := 2*parentIndex + 1
	temp := h[parentIndex]
	for childIndex < length {
		if childIndex+1 < length && h[childIndex] > h[childIndex+1] {
			childIndex = childIndex + 1
		}
		if h[childIndex] > temp {
			break
		}
		h[parentIndex] = h[childIndex]
		parentIndex = childIndex
		childIndex = 2*childIndex + 1
	}
	h[parentIndex] = temp
}
