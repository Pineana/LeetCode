package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	heap := NewHeap(10)
	fmt.Println(heap)
	heap = append(heap, 0)
	upAdjust(heap)
	fmt.Println(heap)
	heap[0], heap[len(heap)-1] = heap[len(heap)-1], heap[0]
	heap = heap[0 : len(heap)-1]
	downAdjust(heap, 0, len(heap))
	fmt.Println(heap)
	sortHeap(heap)
	fmt.Println(heap)
}

type Heap []int

func NewHeap(size int) Heap {

	heap := make([]int, 0)
	rand.Seed(time.Now().Unix())
	for i := 0; i < size; i++ {
		heap = append(heap, rand.Intn(100))
	}
	fmt.Println(heap)
	buildHeap(heap)
	return heap
}

func downAdjust(h Heap, parentIndex int, length int) {
	temp := h[parentIndex]
	childrenIndex := 2*parentIndex + 1
	for childrenIndex < length {
		if childrenIndex+1 < length && h[childrenIndex+1] < h[childrenIndex] {
			childrenIndex++
		}
		if temp < h[childrenIndex] {
			break
		}
		h[parentIndex] = h[childrenIndex]
		parentIndex = childrenIndex
		childrenIndex = 2*childrenIndex + 1
	}
	h[parentIndex] = temp
}

func buildHeap(h Heap) {
	for i := (len(h) - 2) / 2; i >= 0; i-- {
		downAdjust(h, i, len(h))
	}
}

func upAdjust(h Heap) {
	childrenIndex := len(h) - 1
	parentIndex := (childrenIndex - 1) / 2
	temp := h[childrenIndex]
	for childrenIndex != 0 || h[childrenIndex] < h[parentIndex] {
		if temp < h[parentIndex] {
			h[childrenIndex] = h[parentIndex]
			childrenIndex = parentIndex
			parentIndex = (childrenIndex - 1) / 2
		}
	}
	h[childrenIndex] = temp
}

func sortHeap(h Heap) {
	for i := len(h)-1; i > 0; i-- {
		h[0], h[i] = h[i], h[0]
		downAdjust(h, 0, i)
	}
}
