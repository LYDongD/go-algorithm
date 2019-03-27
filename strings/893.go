package main

import "fmt"

func numSpecialEquivGroups(A []string) int {
    if len(A) == 0 {
	return 0
    }

    groupDic := make(map[string]int)
    //分成两部分，分别存放偶数和奇数
    for _, str := range A {
	charCount := make([]rune, 52)
	for index, char := range str {
	    charCount[int(char - 'a') + 26 * ( index % 2)]++
	}
	groupDic[string(charCount)]++
    }

    return len(groupDic)

}

func isSpecialEquiv(A string, B string) bool {

    if A == B {
	return true
    }

    wordDic1 := make(map[rune]int)
    wordDic2 := make(map[rune]int)
    for index, char := range A {
	if index % 2 == 0 {
	    wordDic1[char]++
	}else {
	    wordDic2[char]++
	}
    }

    for index, char := range B {
	if index % 2 == 0 {
	    if wordDic1[char] == 0 {
		return false
	    }
	    wordDic1[char]--
	    if wordDic1[char] == 0 {
		delete(wordDic1, char)
	    }
	}else {
	    if wordDic2[char] == 0 {
		return false
	    }
	    wordDic2[char]--
	    if wordDic2[char] == 0 {
		delete(wordDic2, char)
	    }
	}
    }
    return len(wordDic1) == 0 && len(wordDic2) == 0
}

func main() {
    fmt.Println(numSpecialEquivGroups([]string{"a","b","c","a","c","c"}))
    fmt.Println(numSpecialEquivGroups([]string{"aa","bb","ab","ba"}))
    fmt.Println(numSpecialEquivGroups([]string{"abc","acb","bac","bca","cab","cba"}))
    fmt.Println(numSpecialEquivGroups([]string{"abcd","cdab","adcb","cbad"}))
    fmt.Println(numSpecialEquivGroups([]string{"couxuxaubw","zsptcwcghr","kkntvvhbcc","nkhtcvvckb","crcwhspgzt"}))
    fmt.Println(numSpecialEquivGroups([]string{"abc","acb","bac","bca","cab","cba"}))
}
