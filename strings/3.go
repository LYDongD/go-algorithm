package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	sLength := len(s)
	if sLength <= 1 {
		return sLength
	}

	longest := 1
	start, end := 0, 1
	existed := make(map[byte]int)
	existed[s[start]] = 0
	for end < sLength {

		//update start when repeated
		if repeatedIndex, ok := existed[s[end]]; ok && repeatedIndex >= start {
			start = repeatedIndex + 1
		} else { //cal & update length when no repeated
			diffLen := end - start + 1
			if diffLen > longest {
				longest = diffLen
			}
		}

		//update or add existed char index
		existed[s[end]] = end

		//move end
		end++

		//exit in advance
		if sLength-start+1 <= longest {
			break
		}
	}

	return longest
}

func main() {
	fmt.Println(lengthOfLongestSubstring("tmmzuxt"))
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
