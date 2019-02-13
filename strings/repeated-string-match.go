package main

import "strings"
import "fmt"

func repeatedStringMatch(A string, B string) int {

    n := 1
    S := A
    for len(S) < len(B) {
	S = S + A
	n++
    }

    if strings.Index(S, B) >= 0 {
	return n
    }

    S = S + A
    if strings.Index(S, B) > 0 {
	return n + 1
    }

    return -1
}

func main() {
    fmt.Println(repeatedStringMatch("abc", "cabcabca"))
}
