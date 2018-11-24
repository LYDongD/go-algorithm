//package main
//
//import (
//	"fmt"
//	"math"
//)

func findDisappearedNumbers(nums []int) []int {

	var result []int
	if len(nums) == 0 {
		return result
	}

	//count ele
	for _, num := range nums {
		markIndex := int(math.Abs(float64(num))) - 1
		if nums[markIndex] > 0 {
			nums[markIndex] = -nums[markIndex]
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			result = append(result, i+1)
		}
	}

	return result

}

//func main() {
//
//	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
//	fmt.Println(findDisappearedNumbers(nums))
//}
