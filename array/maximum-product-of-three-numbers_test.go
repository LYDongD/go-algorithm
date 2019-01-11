package array

import "testing"

func Test_maximumProduct(t *testing.T) {
	nums := []int{3, 2, 1}
	result := MaximumProduct(nums)
	t.Log(result)

	nums2 := []int{1, 2, 3, 4}
	t.Log(MaximumProduct(nums2))

	nums3 := []int{-1, -2, -3, 2}
	t.Log(MaximumProduct(nums3))
	// t.Fatal("not implemented")
}
