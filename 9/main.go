package main

func main() {
	isPalindrome(-121)
}

// 解法1：反转整数，然后和原数进行比较
func isPalindrome(x int) bool {
	origin, after := x, 0
	for x != 0 {
		digit := x % 10
		x = x / 10
		after = after*10 + digit
	}
	if origin == after {
		return true
	}
	return false
}

// 解法2：反转后半部分数据，和前半部分比较，需要关注奇位数和偶位数的情况
func isPalindrome2(x int) bool {
	if x < 0 || x%10 == 0 && x != 0 {
		return false
	}
	rev := 0
	for x > rev {
		rev = rev*10 + x%10
		x = x / 10
	}
	return x == rev || x == rev/10
}
