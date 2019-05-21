package main

import "fmt"

func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}

	charDic := []rune{' ', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L',
		'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X',
		'Y', 'Z'}

	}

	duplicatedDic := make(map[string]int)
	return backtrace(charDic, s, "", duplicatedDic, 0)
}


//todo 
func backtrace(charDic []rune, left string, selected string, duplicatedDic map[string]int, count int) int {
	if len(left) == 0 {
		if duplicatedDic[selected] == 0 {
			duplicatedDic[selected]++
			return count + 1
		}
	}

	//select 1
	count = backtrace(charDic, left[1:], selected+charDic[left[0]]), duplicatedDic, count)

	//select 2
	if len(left) >= 2 && string(left[2:]) >= "1" && string(left[2:]) <= "26" {
		count = backtrace(charDic, left[2:], selected+charDic[left[0]]+charDic[left[1]], duplicatedDic, count)
	}

	return count
}

func main() {
	fmt.Println(numDecodings("226"))
}
