package main

import "fmt"

func main() {
	fmt.Println(myAtoi("9223372036854775808"))
}

func myAtoi(s string) int {
	nagitive := false
	max, min := 1<<31-1, -1<<31
	var index, sum int
	for index < len(s) {
		if s[index] != ' ' {
			break
		}
		index++
	}
	if index >= len(s) {
		return 0
	}
	if s[index] == '-' || s[index] == '+' {
		if s[index] == '-' {
			nagitive = true
		}
		index++
	}
	for index < len(s) {
		if s[index]-'0' < 10 && s[index]-'0' >= 0 {
			sum = sum*10 + int(s[index]-'0')
			if sum > max || sum < min {
				if nagitive {
					return min
				} else {
					return max
				}
			}
			index++
		} else {
			break
		}
	}
	if nagitive {
		sum = -sum
	}

	return sum
}
