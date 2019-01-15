package array

import "testing"

func Test_MisMatch(t *testing.T) {
	nums := []int{1, 2, 2, 4}
	t.Log(FindErrorNums(nums))
}
