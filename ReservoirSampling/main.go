package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	s := Constructor(head)
	fmt.Println(s.GetRandom())
	fmt.Println(s.GetRandom())
	fmt.Println(s.GetRandom())
	fmt.Println(s.GetRandom())
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	head *ListNode
}

func Constructor(head *ListNode) Solution {
	return Solution{
		head: head,
	}
}

func (this *Solution) GetRandom() (ans int) {
	var i int
	var node = this.head
	for node != nil {
		i++
		if rand.Intn(i) == 0 {
			ans = node.Val
		}
		node = node.Next
	}
	return ans
}
