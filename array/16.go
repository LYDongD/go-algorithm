package main

import "fmt"
import "sort"

func threeSumClosest(nums []int, target int) int {
    if len(nums) < 3 {
	panic("elements count less than 3")
    }

    sort.Ints(nums)
    sum := 0
    distance := int(^uint(0) >> 1)
    hasExcessedTarget := false
    for i := 0; i <= len(nums) - 3; i++ {
	if i > 0 && nums[i] == nums[i - 1] {
	    continue
	}

	goal := target - nums[i]
	start := i+1
	end := len(nums) - 1

	//剪枝
	if nums[i] + nums[i+1] + nums[i+2] > target {
	    if !hasExcessedTarget {
		hasExcessedTarget = true
	    }else {
		return sum
	    }

	}

	for start < end {
	    if start > i + 1 && nums[start] == nums[start-1] {
		start++
		continue
	    }

	    if end < len(nums) - 1 && nums[end] == nums[end + 1] {
		end--
		continue
	    }

	    if nums[start] + nums[end] == goal {
		return nums[i] + nums[start] + nums[end]
	    }

	    if abs(nums[start] + nums[end] - goal) < distance {
                distance = abs(nums[start] + nums[end] - goal)
                sum = nums[i] + nums[start] + nums[end]
            }

            if nums[start] + nums[end] < goal {
                start++
            }else if nums[start] + nums[end] > goal {
                end--
            }else {
                break
            }

	}
    }

    return sum
}

func abs(a int) int {
    if a >= 0 {
	return a
    }

    return -a
}


func main() {
    fmt.Println(threeSumClosest([]int{1,1,1,1}, 0))
}
