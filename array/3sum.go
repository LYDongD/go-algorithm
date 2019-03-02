package main

import "fmt"
import "sort"

func threeSum(nums []int) [][]int {

    var result [][]int
    if len(nums) < 3 {
	return result
    }

    sort.Ints(nums)
    for i := 0; i < len(nums); i++ {
	if i > 0 && nums[i] == nums[i-1] {
	    continue
	}

	goal := 0 - nums[i]
	triple := []int{}
	start := i + 1
	end := len(nums) - 1

	if start < end && nums[start] + nums[start+1] > goal {
	    break
	}

	for start < end {
	    if start > i + 1 && nums[start] == nums[start - 1] {
		start++
		continue
	    }

	    if end < len(nums) - 1 && nums[end] == nums[end+1] {
		end--
		continue
	    }

	    sum := nums[start] + nums[end]
	    if sum == goal {
		triple = []int{nums[i], nums[start], nums[end]}
		result = append(result, triple)
		start++
		end--
	    }else if sum < goal {
		start++
	    }else {
		end--
	    }
	}

    }
    return result
}

func main() {
    fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}


