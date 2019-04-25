package main

import "fmt"

func commonChars(A []string) []string {
    if len(A) <= 1 {
	return A
    }

    commonDic := make(map[rune]int)
    for _, char := range A[0] {
	commonDic[char]++
    }

    for i := 1; i < len(A); i++ {
	charCount := make(map[rune]int)
	for _, char := range A[i] {
	    charCount[char]++
	}

	for commonChar, count := range commonDic {
	    if count > charCount[commonChar] {
		commonDic[commonChar] = charCount[commonChar]
	    }
	}
    }

    result := []string{}
    for char, count := range commonDic {
	i := count
	for i >= 1 {
	    result = append(result, string(char))
	    i--
	}
    }

    return result
}


func main() {
    fmt.Println(commonChars([]string{"bella","label","roller"}))
    fmt.Println(commonChars([]string{"cool","lock","cook"}))
}
