package main

import "fmt"

func rotateString(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}

	if len(A) == 0 && len(B) == 0 {
		return true
	}

	for i := 0; i < len(A); i++ {
		A := A[i+1:] + A[:i+1]
		if A == B {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(rotateString("abcde", "cdeab"))
	fmt.Println(rotateString("abcde", "abced"))
}
