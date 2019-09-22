package main

import "fmt"

func hIndex(citations []int) int {
    if len(citations) == 0 {
	return 0
    }

    start, end := 0, len(citations) - 1
    for start <= end {
	mid := start + (end - start) / 2
	k := len(citations) - mid
	if citations[mid] >= k && (mid - 1 < 0 || citations[mid-1] <= k) {
	    return k
	}

	if citations[mid] >= k {
	    end = mid - 1
	}else {
	    start = mid + 1
	}
    }

    return 0
}

func main() {
    fmt.Println(hIndex([]int{0,1,3,5,6}))
    fmt.Println(hIndex([]int{1}))
    fmt.Println(hIndex([]int{1,1})
}
