package main

import "fmt"

func checkPossibility(nums []int) bool {

    n := len(nums)
    if n <= 2 {
	return true
    }

    descendNum := 0
    for last, current := 0, 1; current < n; last, current = last+1, current+1 {
	if nums[current] >= nums[last] {
	    continue
	}

	descendNum++
	if descendNum == 2 {
	    return false
	}

	//replace last num
	if last - 1 < 0  || (last - 1) >= 0 && nums[current] >= nums[last-1] {
	    continue
	}

	//replace current num
	if current + 1 == n || (current + 1) < n && nums[current+1] >= nums[last] {
	    nums[current] = nums[last]
	    continue
	}

	return false
    }

    return true
}


func main(){
    nums3 := []int{1,2,3,4,1,3,4}
    fmt.Println(checkPossibility(nums3))
}
