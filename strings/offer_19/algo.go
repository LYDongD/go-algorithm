package main

import "fmt"

func isMatch(s string, p string) bool {
    return match(0, 0, s, p)
}

func match(sIndex,pIndex int, s, p string) bool {

    //fmt.Println(s,p, "s ",sIndex ," p ", pIndex)

    if pIndex >= len(p) && sIndex >= len(s) {
	return true
    }

    if pIndex >= len(p) {
	return false
    }


    char := p[pIndex]
    if pIndex < len(p) - 1 && p[pIndex+1] == '*' {
	mostMatch := count(char, s, sIndex)
	for i := 0; i <= mostMatch; i++ {
	    result :=  match(sIndex + i, pIndex + 2, s, p)
	    if result {
		return true
	    }
	}

	return false
    }

    if sIndex >= len(s) {
	return false
    }

    if char == '.' || char == s[sIndex]  {
	return match(sIndex+1, pIndex+1, s, p)
    }

    return false
}

func count(char byte, s string, current int) int{
    if char == '.' {
	return len(s) - current
    }

    if current >= len(s) || char != s[current] {
	return 0
    }

    count := 1
    next := current + 1
    for next < len(s) && s[next] == s[current]  {
	count++
	current = next
	next++
    }

    return count
}


func main() {
 //   fmt.Println(count('s', "mississippi", 2))
   fmt.Println(isMatch("", ".*"))
   fmt.Println(isMatch("a", "ab*"))
   fmt.Println(isMatch("mississippi", "mis*is*p*."))
   fmt.Println(isMatch("aa", "a"))
   fmt.Println(isMatch("aa", "a*"))
   fmt.Println(isMatch("ab", ".*"))
   fmt.Println(isMatch("aab", "c*a*b"))
   fmt.Println(isMatch("ab", ".*c"))
   fmt.Println(isMatch("aaa", "a*a"))
}
