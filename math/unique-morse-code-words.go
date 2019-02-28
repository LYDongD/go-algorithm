package main

import "fmt"

func uniqueMorseRepresentations(words []string) int {

   morseCodeArr := []string{".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."}

    morseDic := make(map[string]int)
    for _, word := range words {
	morseCode := ""
	for _, char := range word {
	    morseCode = morseCode + morseCodeArr[char - 'a']
	}
	morseDic[morseCode]++
    }

    return len(morseDic)
}

func main() {
    fmt.Println(uniqueMorseRepresentations([]string{"gin", "zen", "gig", "msg"}))
}
