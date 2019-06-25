package main

import "fmt"

var result [][]string

func partition(s string) [][]string {

	result = [][]string{}

	if len(s) == 0 {
		return result
	}

	if len(s) == 1 {
		return append(result, []string{s})
	}

	current := []string{}
	backtrace(s, current)
	return result
}

func backtrace(left string, current []string) {
	if len(left) == 0 {
		result = append(result, current)
		return
	}

	for i := 0; i < len(left); i++ {
		if isPalindrome(left[:i+1]) {
			tmp := make([]string, len(current))
			copy(tmp, current)
			tmp = append(tmp, left[:i+1])
			backtrace(left[i+1:], tmp)
		}
	}

	return
}

func isPalindrome(a string) bool {

	if len(a) == 0 {
		return false
	}

	i, j := 0, len(a)-1
	for i < j {
		if a[i] != a[j] {
			return false
		}

		i++
		j--
	}

	return true
}

func main() {
	fmt.Println(partition("aab"))
}
