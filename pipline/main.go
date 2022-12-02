package main

import (
	"fmt"
	"time"
)

var ch1 = make(chan struct{})
var ch2 = make(chan struct{})
var ch3 = make(chan struct{})

func main() {

	go cat()
	go dog()
	go fish()
	ch1 <- struct{}{}
	time.Sleep(time.Minute)

}

func cat() {
	for {
		select {
		case <-ch1:
			fmt.Println("cat")
			ch2 <- struct{}{}
		}
	}
}

func dog() {
	for {
		select {
		case <-ch2:
			fmt.Println("dog")
			ch3 <- struct{}{}
		}
	}
}

func fish() {
	var i = 0
	for {
		select {
		case <-ch3:
			fmt.Println("fish")
			if i < 9 {
				ch1 <- struct{}{}
				i++
			}
		}
	}
}
