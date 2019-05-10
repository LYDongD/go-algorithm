package main

import "fmt"

func addToArrayForm(A []int, K int) []int {
	if K == 0 || len(A) == 0 {
		return A
	}

	i := len(A) - 1
	carry := 0
	result := []int{}
	for K > 0 || i >= 0 {

		digit := 0
		if K > 0 {
			digit = K % 10
			K = K / 10
		}

		sum := 0
		if i >= 0 {
			sum = A[i] + digit + carry
			i--
		} else {
			sum = digit + carry
		}

		if sum >= 10 {
			sum = sum % 10
			carry = 1
		} else {
			carry = 0
		}

		result = append([]int{sum}, result[0:]...)
	}

	if carry == 1 {
		result = append([]int{1}, result[0:]...)
	}

	return result
}

func main() {
	fmt.Println(addToArrayForm([]int{1, 2, 0, 0}, 34))
	fmt.Println(addToArrayForm([]int{2, 7, 4}, 181))
	fmt.Println(addToArrayForm([]int{2, 1, 5}, 806))
	fmt.Println(addToArrayForm([]int{0}, 23))
	fmt.Println(addToArrayForm([]int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9}, 1))
}
