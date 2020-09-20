package main

import "fmt"

func restoreString(s string, indices []int) string {

	if len(s) <= 1 {
		return s
	}

	resultArr := make([]rune, len(s))
	for i, char := range s {
		resultArr[indices[i]] = char
	}

	return string(resultArr)
}

func main() {
	fmt.Println(restoreString("abc", []int{0, 1, 2}))
	fmt.Println(restoreString("codeleet", []int{4, 5, 6, 7, 0, 2, 1, 3}))
	fmt.Println(restoreString("aaiougrt", []int{4, 0, 2, 6, 7, 3, 1, 5}))
}
