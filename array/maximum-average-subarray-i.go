package array

func findMaxAverage(nums []int, k int) float64 {

	if len(nums) == 0 || k <= 0 {
		panic("illegal parameters")
	}

	var slideSum int
	for i := 0; i < k; i++ {
		slideSum += nums[i]
	}

	maxSum := slideSum
	for i := k; i < len(nums); i++ {
		slideSum = slideSum - nums[i-k] + nums[i]
		if slideSum > maxSum {
			maxSum = slideSum
		}
	}

	return float64(maxSum) / float64(k)
}
