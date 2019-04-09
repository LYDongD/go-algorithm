package main

import "fmt"

func isAlienSorted(words []string, order string) bool {
    if len(words) == 1 {
	return true
    }

    //char index dic
    orderCharDic := make(map[byte]int)
    for index, char := range order {
	orderCharDic[byte(char)] = index
    }

    //compare 2 ele
    for i := 0; i < len(words) - 1; i++ {
	if compare(words[i], words[i+1], orderCharDic) < 0 {
	    return false
	}
    }

    return true
}

func compare(A, B string, orderCharDic map[byte]int) int {
    i := 0
    for i < len(A) && i < len(B) {
	if orderCharDic[(A[i])] > orderCharDic[B[i]] {
	    return -1
	}else if orderCharDic[A[i]] < orderCharDic[B[i]] {
	    return 1
	}

	i++
    }

    if i == len(B) && i < len(A) {
	return -1
    }

    return 1
}

func main() {
    fmt.Println(isAlienSorted([]string{"hello","leetcode"}, "hlabcdefgijkmnopqrstuvwxyz"))
    fmt.Println(isAlienSorted([]string{"word","world","row"}, "worldabcefghijkmnpqstuvxyz"))
    fmt.Println(isAlienSorted([]string{"apple","app"}, "abcdefghijklmnopqrstuvwxyz"))
}
