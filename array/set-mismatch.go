package array

func findErrorNums(nums []int) []int {
	countArr := make([]int, len(nums)+1)
	countArr[0] = -1

	for _, num := range nums {
		countArr[num]++
	}

	result := make([]int, 2)
	for index, num := range countArr {
		if num > 1 {
			result[0] = index
		}

		if num == 0 {
			result[1] = index
		}

		if result[0] > 0 && result[1] > 0 {
			return result
		}
	}

	return result
}
