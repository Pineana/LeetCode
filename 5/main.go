package main

func main() {

}

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	var start, end int
	for i := 0; i < len(s); i++ {
		l1, r1 := expendAroundCenter(s, i, i)
		l2, r2 := expendAroundCenter(s, i, i+1)
		if r1-l1 > end-start {
			start, end = l1, r1
		}
		if r2-l2 > end-start {
			start, end = l2, r2
		}
	}
	return s[start : end+1]
}

func expendAroundCenter(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return left + 1, right - 1
}
