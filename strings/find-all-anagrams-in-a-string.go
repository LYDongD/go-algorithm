package main

import "fmt"

func findAnagrams(s string, p string) []int {

	var startIndexs []int

	if s == "" || p == "" {
		return startIndexs
	}

	for i := range s {

		endIndex := i + len(p)
		if endIndex > len(s) {
			endIndex = len(s)
		}
		subStr := s[i:endIndex]
		if isAnagrams(subStr, p) {
			startIndexs = append(startIndexs, i)
		}
	}

	return startIndexs
}

func isAnagrams(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	checkList := make([]int, 26)
	for _, c := range s1 {
		checkList[c-'a']++
	}

	for _, c := range s2 {
		if checkList[c-'a'] == 0 {
			return false
		}
		checkList[c-'a']--
	}

	return true

}

//func main() {
//
//	s := "cbaebabacd"
//	p := "abc"
//	fmt.Println(findAnagrams(s, p))
//
//	s = "abab"
//	p = "ab"
//	fmt.Println(findAnagrams(s, p))
//
//}
