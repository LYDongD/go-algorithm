package array

import "testing"

func Test_FindMaxAverage(t *testing.T) {
	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4
	t.Log(FindMaxAverage(nums, k))
}
