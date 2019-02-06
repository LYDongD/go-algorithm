package main

import "fmt"

func toLowerCase(str string) string {
	if len(str) == 0 {
		return str
	}

	result := ""
	delta := 'a' - 'A'
	for _, char := range str {
		if char >= 'A' && char <= 'Z' {
			result = result + string(char+delta)
		} else {
			result = result + string(char)
		}
	}

	return result
}

func main() {
	fmt.Println(toLowerCase("Hello"))
	fmt.Println(toLowerCase("here"))
	fmt.Println(toLowerCase("LOVELY"))
}
