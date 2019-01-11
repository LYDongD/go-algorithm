package array

//import "fmt"
import "sort"

func maximumProduct(nums []int) int {
	if len(nums) < 3 {
		panic("need at least 3 num")
	}

	if len(nums) == 3 {
		return nums[0] * nums[1] * nums[2]
	}

	//quickSort(nums, 0, len(nums)-1)
	sort.Ints(nums)

	n := len(nums)
	product1 := nums[n-1] * nums[0] * nums[1]
	product2 := nums[n-1] * nums[n-2] * nums[n-3]
	if product1 > product2 {
		return product1
	} else {
		return product2
	}

}

func quickSort(nums []int, start int, end int) {
	if start >= end {
		return
	}

	q := partion(nums, start, end)
	quickSort(nums, start, q-1)
	quickSort(nums, q+1, end)
}

func partion(nums []int, start int, end int) int {
	pivort := nums[end]
	q := start
	for i := start; i < end; i++ {
		if nums[i] < pivort {
			swap(nums, q, i)
			q++
		}
	}

	swap(nums, q, end)

	return q
}

func swap(nums []int, a int, b int) {
	nums[a], nums[b] = nums[b], nums[a]
}
