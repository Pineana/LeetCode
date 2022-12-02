package main

import "fmt"

func main() {
	fmt.Println(jumpFloor(7))
	fmt.Println(jumpFloorN(7))
}

func jumpFloor(n int) int {
	if n >= 0 && n <= 2 {
		return n
	}
	return jumpFloor(n-1) + jumpFloor(n-2)
}

func jumpFloorN(n int) int {
	if n >= 0 && n <= 2 {
		return n
	}
	return jumpFloorN(n-1) * 2
}
