package go_algorithm

import (
	"fmt"
	"strconv"
)

func addStrings(num1 string, num2 string) string {
	var maxLength int
	num1Len := len(num1)
	num2Len := len(num2)
	if num1Len < num2Len {
		maxLength = num2Len
	} else {
		maxLength = num1Len
	}

	var carry bool
	var result string
	for i := 0; i < maxLength; i++ {
		sum := 0
		if i < num1Len {
			j, err := strconv.Atoi(num1[num1Len-i-1 : num1Len-i])
			if err == nil {
				sum += j
			}
		}

		if i < num2Len {
			j, err := strconv.Atoi(num2[num2Len-i-1 : num2Len-i])
			if err == nil {
				sum += j
			}
		}

		if carry {
			sum += 1
		}

		if sum >= 10 {
			carry = true
			sum = sum % 10
		} else {
			carry = false
		}

		fmt.Println("sum: ", sum, "carry: ", carry)
		result = strconv.Itoa(sum) + result
	}

	if carry {
		result = "1" + result
	}

	return result
}

//func main() {
//	fmt.Println(addStrings("123456789", "987654321"))
//}
