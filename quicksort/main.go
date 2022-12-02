package main

import (
	"fmt"
	"math/rand"
	"time"
)

type stack struct {
	data []int
	len  int
}

func NewStack() *stack {
	return &stack{
		data: make([]int, 0),
	}
}

func (s *stack) pop() int {
	if s.len > 0 {
		popItem := s.data[s.len-1]
		s.len--
		s.data = s.data[0:s.len]
		return popItem
	} else {
		panic("stack is empty")
	}
}

func (s *stack) push(item int) {
	s.data = append(s.data, item)
	s.len++
}

func main() {
	rl := buildSlice(10)
	rl2 := buildSlice(10)
	fmt.Println(rl)
	quickSort(rl, 0, len(rl)-1)
	fmt.Println(rl)
	fmt.Println(rl2)
	quickSortWithStack(rl2, 0, len(rl2)-1)
	fmt.Println(rl2)
}

func buildSlice(len int) []int {
	rand.Seed(time.Now().Unix())
	list := make([]int, 0)
	for i := 0; i < len; i++ {
		list = append(list, rand.Intn(100))
	}
	return list
}

func quickSort(rl []int, L int, R int) {
	if L >= R {
		return
	}
	pivotIndex := partitionDouble(rl, L, R)
	quickSort(rl, L, pivotIndex-1)
	quickSort(rl, pivotIndex+1, R)
}

func quickSortWithStack(rl []int, L, R int) {
	if L >= R {
		return
	}
	s := NewStack()
	pivotIndex := partitionSingle(rl, L, R)
	if pivotIndex-1 > L {
		s.push(L)
		s.push(pivotIndex - 1)
	}
	if pivotIndex+1 < R {
		s.push(pivotIndex + 1)
		s.push(R)
	}
	for s.len != 0 {
		R = s.pop()
		L = s.pop()
		pivotIndex := partitionSingle(rl, L, R)
		if pivotIndex-1 > L {
			s.push(L)
			s.push(pivotIndex - 1)
		}
		if pivotIndex+1 < R {
			s.push(pivotIndex + 1)
			s.push(R)
		}
	}
}

func partitionDouble(rl []int, L int, R int) int {

	mid := (L + R) / 2
	rl[L], rl[mid] = rl[L], rl[mid]
	pivot := L
	for L < R {
		for rl[R] >= rl[pivot] && L < R {
			R--
		}
		for rl[L] <= rl[pivot] && L < R {
			L++
		}
		rl[R], rl[L] = rl[L], rl[R]
		if L >= R {
			rl[L], rl[pivot] = rl[pivot], rl[L]
		}
	}
	return L
}

func partitionSingle(rl []int, L int, R int) int {
	mark := L
	for i := L; i <= R; i++ {
		if rl[i] < rl[L] {
			mark++
			rl[i], rl[mark] = rl[mark], rl[i]
		}
	}
	rl[mark], rl[L] = rl[L], rl[mark]
	return mark
}
