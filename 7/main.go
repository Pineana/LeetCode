package main

import (
	"fmt"
	"math"
)

func main() {
	reverse(1563847412)
}

func reverse(x int) int {
	var rev int
	Max,Min := (1<<31)-1,-(1<<31)
	fmt.Println(Max,Min)
	for x != 0 {
		if rev > math.MaxInt32 || rev < -math.MaxInt32 {
			return 0
		}
		digit := x % 10
		x = x / 10
		rev = rev*10 + digit
	}
	return rev
}
