package main

import "fmt"

func makeGood(s string) string {
	if len(s) <= 1 {
		return s
	}

	i := 0
	for i <= len(s)-2 {
		if lowerUpperCase(s[i], s[i+1]) {
			s = s[0:i] + s[i+2:]
			i = i - 1
			if i < 0 {
				i = 0
			}
		} else {
			i++
		}
	}

	return s
}

func lowerUpperCase(a, b byte) bool {
	return a+32 == b || b+32 == a
}

func main() {
	fmt.Println(makeGood("leEeetcode"))
	fmt.Println(makeGood("abBAcC"))
}
