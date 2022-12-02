package main

import "fmt"

func main() {
	lru := NewLRUCache(4)
	lru.Put("a", "a")
	lru.Put("a", "a1")
	lru.Put("b", "b")
	lru.Put("c", "c")
	lru.Put("d", "d")
	fmt.Println(lru.Get("a"))
	fmt.Println(lru.Get("b"))
	fmt.Println(lru.Get("c"))
	fmt.Println(lru.Get("d"))
	fmt.Println(lru.Get("e"))
}

type LinkList struct {
	Head *LinkNode
	Tail *LinkNode
}

type LinkNode struct {
	Key   string
	Value string
	Pre   *LinkNode
	Next  *LinkNode
}

func (l *LinkList) Insert(node *LinkNode) {
	node.Pre = l.Head
	node.Next = l.Head.Next
	l.Head.Next.Pre = node
	l.Head.Next = node
}

func (l *LinkList) Delete(node *LinkNode) {
	pre := node.Pre
	next := node.Next
	pre.Next = next
	next.Pre = pre
}

func (l *LinkList) PopTail() *LinkNode {
	item := l.Tail.Pre
	l.Tail.Pre = item.Pre
	item.Pre.Next = l.Tail
	return item
}

func NewLinkList() *LinkList {
	head := &LinkNode{}
	tail := &LinkNode{}
	head.Next = tail
	tail.Pre = head
	return &LinkList{
		Head: head,
		Tail: tail,
	}
}

type LRUCache struct {
	HashTable map[string]*LinkNode
	Data      *LinkList
	Cap       int
}

func (l *LRUCache) Put(key string, value string) {
	if v, ok := l.HashTable[key]; ok {
		v.Value = value
		l.Data.Delete(v)
		l.Data.Insert(v)
	} else {
		node := &LinkNode{
			Key:   key,
			Value: value,
		}
		if len(l.HashTable) == l.Cap {
			popItem := l.Data.PopTail()
			delete(l.HashTable, popItem.Key)
		}
		l.HashTable[key] = node
		l.Data.Insert(node)
	}
}

func (l *LRUCache) Get(key string) string {
	if v, ok := l.HashTable[key]; ok {
		return v.Value
	} else {
		return "nil"
	}
}

func NewLRUCache(n int) *LRUCache {
	return &LRUCache{
		HashTable: make(map[string]*LinkNode, n),
		Data:      NewLinkList(),
		Cap:       n,
	}
}
