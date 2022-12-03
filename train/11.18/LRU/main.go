package main

import (
	"container/list"
	"fmt"
)

func main() {
	lru := NewLRU(3)
	lru.Put("a", "a")
	fmt.Println(lru.Get("a"))
	lru.Put("a", "a1")
	fmt.Println(lru.Get("a"))
	lru.Put("b", "b")
	fmt.Println(lru.Get("a"))
	lru.Put("c", "c")
	lru.Put("d", "d")
	fmt.Println(lru.Get("a"))
	fmt.Println(lru.Get("b"))

}

type LRU struct {
	Data    *list.List
	HashMap map[string]*list.Element
	Cap     int
}

type Node struct {
	key   string
	value string
}

func NewLRU(n int) *LRU {
	return &LRU{
		Data:    list.New(),
		HashMap: make(map[string]*list.Element, 0),
		Cap:     n,
	}
}

func (l *LRU) Put(key string, value string) {
	node := &Node{
		key:   key,
		value: value,
	}
	if v, ok := l.HashMap[key]; !ok {
		if len(l.HashMap) >= l.Cap {
			rItem := l.Data.Remove(l.Data.Back())
			delete(l.HashMap, rItem.(*Node).key)
			item := l.Data.PushFront(node)
			l.HashMap[key] = item
		} else {
			item := l.Data.PushFront(node)
			l.HashMap[key] = item
		}
	} else {
		l.Data.MoveToFront(v)
		v.Value.(*Node).value = value
	}
}

func (l *LRU) Get(key string) string {
	if v, ok := l.HashMap[key]; ok {
		l.Data.MoveToFront(v)
		return v.Value.(*Node).value
	} else {
		return "nil"
	}
}
