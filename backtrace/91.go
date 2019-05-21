package main

import "fmt"
import "strconv"

func numDecodings(s string) int {
	if len(s) == 0 {
		return 0
	}

	charDic := []rune{' ', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L',
		'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X',
		'Y', 'Z'}

	duplicatedDic := make(map[string]int)
	return backtrace(charDic, s, "", duplicatedDic, 0)
}

func backtrace(charDic []rune, left string, selected string, duplicatedDic map[string]int, count int) int {
	if len(left) == 0 {
		fmt.Println("count: ", count, "selected: ", selected)
		if duplicatedDic[selected] == 0 {
			duplicatedDic[selected]++
			return count + 1
		} else {
			return count
		}
	}

	//select 1
	first, err := strconv.Atoi(left[0:1])
	if err == nil && first >= 1 && first <= 26 {
		count = backtrace(charDic, left[1:], selected+string(charDic[first]), duplicatedDic, count)
	}

	//select 2
	if len(left) >= 2 && left[0:1] != "0" {
		double, err1 := strconv.Atoi(left[0:2])
		if double >= 1 && double <= 26 {
			fmt.Println("select 2", selected+string(charDic[double]))
			if err1 == nil {
				count = backtrace(charDic, left[2:], selected+string(charDic[double]), duplicatedDic, count)
			}

		}
	}

	return count
}

func main() {
	fmt.Println(numDecodings("226"))
	fmt.Println(numDecodings("12"))
	fmt.Println(numDecodings("0"))
	fmt.Println(numDecodings("01"))
}
