package main

import "strconv"
import "fmt"

func largestNumber(nums []int) string {
	if len(nums) == 0 {
		return ""
	}

	if len(nums) == 1 {
		return strconv.Itoa(nums[0])
	}

	searchedDic := make(map[int]int)
	result := ""
	for len(searchedDic) < len(nums) {
		maxStr := ""
		maxStrIndex := 0
		for index, num := range nums {
			if searchedDic[index] > 0 {
				continue
			}

			maxStr = compare(maxStr, strconv.Itoa(num))
			if maxStr == strconv.Itoa(num) {
				maxStrIndex = index
			}
		}

		if len(result) == 0 && maxStr == "0" {
			return "0"
		}

		searchedDic[maxStrIndex]++
		result = result + maxStr
	}

	return result
}

func compare(a, b string) string {
	if a+b > b+a {
		return a
	}

	return b
}

func main() {
	fmt.Println(largestNumber([]int{10, 2}))
	fmt.Println(largestNumber([]int{3, 30, 34, 5, 9}))
	fmt.Println(largestNumber([]int{0, 0}))
}
