package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	fmt.Println(ReserviorSampling())
	fmt.Println(ReserviorSampling())
	fmt.Println(ReserviorSampling())
	fmt.Println(ReserviorSampling())
	fmt.Println(ReserviorSampling())
}

func ReserviorSampling() int {
	var i int = 1
	for rand.Intn(i) == 0 {
		i++
	}
	return i
}
