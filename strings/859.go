package main

import "fmt"

func buddyStrings(A string, B string) bool {
    if len(A) != len(B) {
	return false
    }

    differs := []byte{}
    hasRepeated := false
    wordCount := make(map[byte]int)
    for i := 0; i < len(A); i++ {
	if A[i] != B[i] {
	    differs = append(differs, A[i])
	    differs = append(differs, B[i])
	}

	if wordCount[A[i]] > 0 {
	    hasRepeated = true
	}else {
	    wordCount[A[i]]++
	}
    }

    if len(differs) > 0 {
	if len(differs) != 4 {
	    return false
	}

	if differs[0] == differs[3] && differs[1] == differs[2] {
	    return true
	}

	return false
    }

    return hasRepeated

}

func main() {
    fmt.Println(buddyStrings("aaaaaaabc","aaaaaaacb"))
    fmt.Println(buddyStrings("aa","aa"))
    fmt.Println(buddyStrings("ab","ba"))
    fmt.Println(buddyStrings("","aa"))
    fmt.Println(buddyStrings("abab","abab"))
}
