package array

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max := 0
	count := 0
	for _, num := range nums {
		if num == 1 {
			count++
		} else {
			if max == 0 || count > max {
				max = count
			}
			count = 0
		}
	}

	if count > 0 && count > max {
		max = count
	}

	return max
}

func array() {

	nums1 := []int{1, 1, 0, 1, 1, 1}
	nums2 := []int{1, 0, 1, 1, 0, 1}
	nums3 := []int{1, 1, 1}
	fmt.Println(findMaxConsecutiveOnes(nums1))
	fmt.Println(findMaxConsecutiveOnes(nums2))
	fmt.Println(findMaxConsecutiveOnes(nums3))

}
