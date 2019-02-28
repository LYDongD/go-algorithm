package main

import "fmt"

func makesquare(nums []int) bool {
	if len(nums) < 4 {
		return false
	}

	sum := 0
	for _, num := range nums {
		sum += num
	}

	//square must be 4xsideLength
	if sum%4 != 0 {
		return false
	}

	sides := make([]int, 4)
	return dfs(nums, sides, 0, sum/4)
}

func dfs(nums []int, sides []int, index int, sideLength int) bool {

	if index == len(nums) {
		return sides[0] == sides[1] && sides[1] == sides[2] && sides[2] == sides[3]
	}

	//take num to each sides
	for i := 0; i < len(sides); i++ {
		if sides[i]+nums[index] > sideLength {
			continue
		}

		sides[i] = sides[i] + nums[index]
		if dfs(nums, sides, index+1, sideLength) {
			return true
		}

		sides[i] = sides[i] - nums[index]
	}

	return false
}

func main() {
	fmt.Println(makesquare([]int{1, 1, 2, 2, 2}))
	fmt.Println(makesquare([]int{3, 3, 3, 3, 4}))
	fmt.Println(makesquare([]int{5, 5, 5, 5, 4, 4, 4, 4, 3, 3, 3, 3}))
}
