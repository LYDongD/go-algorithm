package main

import "fmt"

func reverseStr(s string, k int) string {
	if len(s) == 0 || k < 0 {
		return ""
	}

	if k == 0 {
		return s
	}

	n := len(s)
	chars := []rune(s)
	cycle := 2 * k
	groups := n / cycle
	left := n % cycle

	for i := 0; i < groups; i++ {
		reverse(chars, i*cycle, i*cycle+k-1)
	}

	var endIndex int
	if left <= k {
		endIndex = n - 1
	} else {
		endIndex = groups*cycle + k - 1
	}

	reverse(chars, groups*cycle, endIndex)

	return string(chars)

}

func reverse(chars []rune, start int, end int) {

	for start < end {
		chars[start], chars[end] = chars[end], chars[start]
		start++
		end--
	}
}

func main() {
	str := "abcdefg"
	reverseStr(str, 2)
	fmt.Println(reverseStr(str, 2))
}
