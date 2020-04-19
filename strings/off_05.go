package main

import "fmt"

func replaceSpace(s string) string {
	if len(s) == 0 {
		return s
	}

	result := ""
	for _, char := range s {
		if char == ' ' {
			result += "%20"
		} else {
			result += string(char)
		}
	}

	return result
}

func main() {
	fmt.Println(replaceSpace("We are happy."))
}
