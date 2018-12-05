package main

import (
	"fmt"
	"math"
	"strconv"
)

func largestPalindrome(n int) int {

	if n == 1 {
		return 9
	}

	start, end := findMultiplierRange(n)

	isFound := false
	maxProduct := end * end
	firstHalf := maxProduct / int(math.Pow10(n))
	assumeLargestPalindrome := 0
	for isFound == false {
		assumeLargestPalindrome = makePalindrome(firstHalf, n)

		for i := end; i >= start; i-- {

			//1 assumeLargestPalindrome may larger than maxProduct
			//2 if i *  j -> i * i -> j * i
			if assumeLargestPalindrome > maxProduct || i*i < assumeLargestPalindrome {
				break
			}

			if assumeLargestPalindrome%i == 0 {
				isFound = true
				break
			}
		}

		//update assume largest palindrome
		firstHalf--
	}

	return assumeLargestPalindrome % 1337

}

func findMultiplierRange(n int) (int, int) {
	return int(math.Pow10(n - 1)), int(math.Pow10(n)) - 1
}

func makePalindrome(firstHalf int, n int) int {
	firstHalfStr := strconv.Itoa(firstHalf)
	palindromeStr := firstHalfStr + reverseStr(firstHalfStr)
	result, _ := strconv.Atoi(palindromeStr)
	return result
}

func reverseStr(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

func main() {
	fmt.Println(largestPalindrome(4))
	fmt.Println(largestPalindrome(3))
	fmt.Println(largestPalindrome(2))
	fmt.Println(largestPalindrome(1))

}
