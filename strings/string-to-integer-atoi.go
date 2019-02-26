package main

import "fmt"

func myAtoi(str string) int {
	//remove white space
	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			str = str[i:]
			break
		}
	}

	if len(str) == 0 {
		return 0
	}

	//first char check
	firstChar := str[0]
	if !isNumChar(firstChar) && !isSignChar(firstChar) {
		return 0
	}

	signFlag := 0
	if isSignChar(firstChar) {
		if len(str) == 1 {
			return 0
		}

		if !isNumChar(str[1]) {
			return 0
		}

		if firstChar == '-' {
			signFlag = -1
		} else if firstChar == '+' {
			signFlag = 1
		}
	}

	start := 0
	if signFlag != 0 {
		start = 1
	}

	result := 0
	INT_MAX := 1<<31 - 1
	INT_MIN := -(1 << 31)
	for i := start; i < len(str); i++ {
		if !isNumChar(str[i]) {
			break
		}

		//overflow
		result = result*10 + int(str[i]-'0')
		if signFlag >= 0 && result > INT_MAX {
			return INT_MAX
		}

		if signFlag < 0 && result*signFlag < INT_MIN {
			return INT_MIN
		}
	}

	//handdle sign
	if signFlag != 0 {
		result = result * signFlag
	}

	return result
}

func isNumChar(char byte) bool {
	return char >= '0' && char <= '9'
}

func isSignChar(char byte) bool {
	return char == '+' || char == '-'
}

func main() {
	fmt.Println(myAtoi("42"))
	fmt.Println(myAtoi("-42"))
	fmt.Println(myAtoi("4193 with words"))
	fmt.Println(myAtoi("words and 987"))
	fmt.Println(myAtoi("-2147483649"))
	fmt.Println(myAtoi("0000256"))
	fmt.Println(myAtoi("   +O 123"))
	fmt.Println(myAtoi("-000000000000001"))
	fmt.Println(myAtoi("9223372036854775808"))
}
