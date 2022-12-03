package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) (head *ListNode) {
	var carry int
	var tail *ListNode
	for l1 != nil || l2 != nil {
		var n1, n2 int
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := (n1 + n2 + carry) % 10
		carry = (n1 + n2 + carry) / 10
		if head == nil {
			head = &ListNode{
				Val: sum,
			}
			tail = head
		} else {
			tail.Next = &ListNode{
				Val: sum,
			}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{
			Val: carry,
		}
	}
	return head
}

func lengthOfLongestSubstring(s string) int {
	var i, j, ans int
	bt := []byte(s)
	mp := make(map[byte]int)
	for ; j < len(s); j++ {
		i = max(i, mp[bt[j]])
		mp[bt[j]] = j + 1
		ans = max(ans, j-i+1)
	}
	return ans
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	var start, end int
	for i := 0; i < len(s); i++ {
		l1, r1 := expend(s, i, i)
		l2, r2 := expend(s, i, i+1)
		if r1-l1 > end-start {
			start, end = l1, r1
		}
		if r2-l2 > end-start {
			start, end = l2, r2
		}
	}
	return s[start : end+1]
}

func expend(s string, left, right int) (int, int) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return left + 1, right - 1
}

func reverse(x int) int {
	var max, min = 1<<31 - 1, -1 << 31
	var ans int
	for x != 0 {
		if ans > max/10 || ans < min/10 {
			return -1
		}
		digit := x % 10
		x = x / 10
		ans = ans*10 + digit
	}
	return ans
}
