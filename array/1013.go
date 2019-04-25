package main

import "fmt"

func canThreePartsEqualSum(A []int) bool {
    if len(A) < 3 {
	return false
    }

    //cal sum
    sum := 0
    for _, num := range A {
	sum += num
    }

    if sum % 3 != 0 {
	return false
    }

    partSum := sum / 3
    partOneSum := 0
    for i := 0; i < len(A)-2; i++ {
	partOneSum += A[i]
	if partOneSum == partSum {
	    partTwoSum := 0
	    for j := i + 1; j < len(A) - 1; j++ {
		partTwoSum += A[j]
		if partTwoSum == partSum {
		    partTreeSum := 0
		    for k := j + 1; k < len(A); k++ {
			partTreeSum += A[k]
		    }

		    if partTreeSum == partSum {
			return true
		    }
		}
	    }
	}
    }

    return false
}

func main() {
    fmt.Println(canThreePartsEqualSum([]int{0,2,1,-6,6,-7,9,1,2,0,1}))
    fmt.Println(canThreePartsEqualSum([]int{0,2,1,-6,6,7,9,-1,2,0,1}))
    fmt.Println(canThreePartsEqualSum([]int{3,3,6,5,-2,2,5,1,-9,4}))
}
