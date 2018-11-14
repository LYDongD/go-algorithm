package main

import "fmt"

func thirdMax(nums []int) int {
	
	const INT_MIN = ^(int(^uint(0) >> 1))
	
	max1st, max2nd, max3rd := INT_MIN, INT_MIN, INT_MIN
	for _, num := range nums {
		if num >= max1st {
			if num != max1st {
				max3rd = max2nd
				max2nd = max1st
				max1st = num
			}
		}else if num >= max2nd {
			if num != max2nd {
				max3rd = max2nd
				max2nd = num
			}
		}else if num > max3rd {
			max3rd = num
		}
	}

	if max3rd != INT_MIN {
		return max3rd
	}else {
		return max1st
	}

}


func main() {

	nums := []int{2,2,3,1}
	fmt.Println(thirdMax(nums))
}

