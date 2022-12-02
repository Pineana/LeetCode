package main

func main() {

}

func reverseWords(s string) string {
	var quick, slow int
	bl := []byte(s)
	for quick < len(bl) && bl[quick] == ' ' {
		quick++
	}
	for ; quick < len(bl); quick++ {
		if quick-1 > 0 && bl[quick-1] == bl[quick] && bl[quick] == ' ' {
			continue
		}
		bl[slow] = bl[quick]
		slow++
	}
	if slow-1 > 0 && bl[slow-1] == ' ' {
		bl = bl[:slow-1]
	} else {
		bl = bl[:slow]
	}
	reverse(bl, 0, len(bl)-1)
	for i := 0; i < len(bl); i++ {
		j := i
		for ; j < len(bl) && bl[j] != ' '; j++ {
		}
		reverse(bl, i, j-1)
		i = j
	}
	return string(bl)
}

func reverse(bl []byte, start, end int) {
	for start < end {
		bl[start], bl[end] = bl[end], bl[start]
		start++
		end--
	}
}
