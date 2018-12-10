package main

import (
	"fmt"
	"strconv"
)

func convertToBase7(num int) string {

	if num == 0 {
		return "0"
	}

	isNegative := false
	if num < 0 {
		num = -num
		isNegative = true
	}
	result := ""
	for num > 0 {
		bit := num % 7
		result = strconv.Itoa(bit) + result
		num = num / 7
	}

	if isNegative {
		result = "-" + result
	}

	return result
}

func main() {
	fmt.Println(convertToBase7(100))
	fmt.Println(convertToBase7(-7))
}
