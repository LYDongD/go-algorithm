package main

import "fmt"

func uncommonFromSentences(A string, B string) []string {
    wordCount := make(map[string]int)

    S := A + " " +  B + " "
    start, end := 0, 0
    for i := 0; i < len(S); i++ {
	if S[i] == ' '{
	    end = i
	    word := S[start:end]
	    wordCount[word]++
	    start = i + 1
	}
    }


    result := []string{}
    for word, count := range wordCount {
	if count == 1 {
	    result = append(result, word)
	}
    }

    return result
}

func main() {
    fmt.Println(uncommonFromSentences("this apple is sweet", "this apple is sour"))
    fmt.Println(uncommonFromSentences("apple apple", "banana"))
}
