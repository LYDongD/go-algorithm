package main

import "fmt"
import "sort"

func hIndex(citations []int) int {
	if len(citations) == 0 {
		return 0
	}

	hIndex := 0
	sort.Ints(citations)
	for i := len(citations) - 1; i >= 0; i-- {
		count := len(citations) - i
		if citations[i] < count {
			break
		}

		if i-1 < 0 || citations[i-1] <= count {
			hIndex = count
			break
		}
	}

	return hIndex
}

func main() {
	fmt.Println(hIndex([]int{3, 0, 6, 1, 5}))
	fmt.Println(hIndex([]int{1, 1}))
}
