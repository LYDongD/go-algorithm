//package main
//
//import (
//	"fmt"
//)

func repeatedSubstringPattern(s string) bool {

	if len(s) <= 1 {
		return false
	}

	for i := len(s) / 2; i >= 1; i-- {

		subStr := findSubStr(s, i)
		if len(subStr) > 0 {
			if copySubStr(subStr, len(s)/len(subStr)) == s {
				return true
			}
		}
	}

	return false
}

func findSubStr(s string, subStrLength int) string {
	if len(s)%subStrLength != 0 {
		return ""
	}

	return s[0:subStrLength]
}

func copySubStr(subStr string, times int) string {

	result := ""
	for times > 0 {
		result = result + subStr
		times--
	}

	return result
}

//func main() {
//
//	s := "bb"
//	fmt.Println(repeatedSubstringPattern(s))
//
//}
