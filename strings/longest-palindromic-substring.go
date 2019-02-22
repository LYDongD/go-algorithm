package main

import "fmt"

func longestPalindrome(s string) string {
    if len(s) <= 1 {
	return s
    }

    //填充*，合并偶数或奇数长的回文串为同一种情况
    filledStr := ""
    for _, char := range s {
	filledStr = filledStr + "*" + string(char)
    }

    filledStr = filledStr + "*"

    //从中间字符开始判断回文
    maxLength := 0
    maxLengthStr := ""
    for m := 1; m < len(filledStr) - 1; m++ {
	l,r := m-1, m+1
	for l >=0 && r <= len(filledStr) - 1 {
	    if filledStr[l] != filledStr[r] {
		break
	    }
	    l--
	    r++
	}

	//算长度时减去*的占位
	length := r - l - 2
	if length > maxLength {
	    maxLength = length
	    maxLengthStr = filledStr[l+1:r]
	}
    }

    //去掉*
    result := ""
    for _, char := range maxLengthStr {
	if char != '*' {
	    result = result + string(char)
	}
    }

    return result
}

func main() {
    fmt.Println(longestPalindrome("babad"))
    fmt.Println(longestPalindrome("cbbd"))
}
