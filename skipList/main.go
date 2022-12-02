package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	fmt.Println(RandLevel())
	fmt.Println(RandLevel())
	fmt.Println(RandLevel())
	fmt.Println(RandLevel())
	fmt.Println(RandLevel())
	fmt.Println(RandLevel())
	fmt.Println(RandLevel())
	RandLevel()
	skipList := NewSkipList()
	skipList.AddElement(&Element{Key: "a", Value: "a"}, 1)
	skipList.AddElement(&Element{Key: "b", Value: "b"}, 2)
	skipList.AddElement(&Element{Key: "c", Value: "c"}, 3)
	skipList.AddElement(&Element{Key: "d", Value: "d"}, 4)
	skipList.AddElement(&Element{Key: "e", Value: "e"}, 5)
}

const MAXSKIPLISTLEVEL = 64

type Element struct {
	Key   string
	Value string
}

type SkipListNode struct {
	Element *Element
	Score   int
	Levels  []*SkipListLevel
	Pre     *SkipListNode
}

type SkipListLevel struct {
	Span int
	Next *SkipListNode
}

type SkipList struct {
	Head, Tail *SkipListNode
	Level      int
}

func NewSkipList() *SkipList {
	s := &SkipList{
		Head: &SkipListNode{
			Levels: make([]*SkipListLevel, MAXSKIPLISTLEVEL),
		},
		Tail: &SkipListNode{
			Levels: make([]*SkipListLevel, MAXSKIPLISTLEVEL),
		},
	}
	for i := 0; i < MAXSKIPLISTLEVEL; i++ {
		s.Head.Levels[i] = &SkipListLevel{
			Span: 0,
		}
		s.Tail.Levels[i] = &SkipListLevel{
			Span: 0,
		}
	}
	return s
}

func (s *SkipList) AddElement(element *Element, score int) {

	newNode := NewSkipListNode(element, score)
	level := len(newNode.Levels)
	var preNodeList = make([]*SkipListNode, level)
	for i := s.Level - 1; i >= 0; i-- {
		Node := s.Head
		for Node.Levels[i].Next != nil {
			if Node.Levels[i].Next.Score > score {
				break
			}
			Node = Node.Levels[i].Next
		}
		preNodeList[i] = Node
	}
	if level > s.Level {
		for i := s.Level; i < level; i++ {
			preNodeList[i] = s.Head
		}
		s.Level = level
	}

	for i, v := range preNodeList {
		newNode.Levels[i].Next = v.Levels[i].Next
		v.Levels[i].Next = newNode
	}
	if preNodeList[0] != s.Head {
		newNode.Pre = preNodeList[0]
	}
}

func (s *SkipList) RemoveElement(element *Element, score int) {

}

func (s *SkipList) GetElement(key string) *Element {

	return &Element{}
}

func RandLevel() int {
	var i int = 1
	for {
		if rand.Intn(4) == 0 {
			break
		}
		i++
	}
	if i > MAXSKIPLISTLEVEL {
		return MAXSKIPLISTLEVEL
	}
	return i
}

func NewSkipListNode(element *Element, score int) *SkipListNode {
	level := RandLevel()
	n := &SkipListNode{
		Element: element,
		Score:   score,
		Levels:  make([]*SkipListLevel, level),
	}
	for i := 0; i < level; i++ {
		n.Levels[i] = &SkipListLevel{}
	}
	return n
}
