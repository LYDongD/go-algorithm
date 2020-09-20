package main

import "fmt"

func findLatestStep(arr []int, m int) int {
	if len(arr) == 0 || m <= 0 {
		return -1
	}

	bits := make([]int, len(arr))
	lastStep := -1
	for i := 0; i < len(arr); i++ {
		bits[arr[i]-1] = bits[arr[i]-1] + 1
		groupLen := 0
		for k := 0; k < len(bits); k++ {
			if bits[k] == 1 {
				groupLen++
			} else {
				if groupLen == m {
					lastStep = i + 1
				}

				groupLen = 0
			}

		}

		if groupLen == m {
			lastStep = i + 1
		}

	}

	return lastStep
}

func main() {
	fmt.Println(findLatestStep([]int{1, 2}, 1))
	fmt.Println(findLatestStep([]int{3, 5, 1, 2, 4}, 1))
	fmt.Println(findLatestStep([]int{3, 1, 5, 4, 2}, 2))
	fmt.Println(findLatestStep([]int{2, 1}, 2))
	fmt.Println(findLatestStep([]int{1}, 1))
}
