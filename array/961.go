package main

import "fmt"

func repeatedNTimes(A []int) int {
    if len(A) == 0 {
	return -1
    }

    dic := make(map[int]int)
    for _, num := range A {
	if dic[num] > 0 {
	    return num
	}

	dic[num]++
    }

    return -1
}

func main() {
    fmt.Println(repeatedNTimes([]int{1,2,3,3}))
    fmt.Println(repeatedNTimes([]int{2,1,2,5,3,2}))
    fmt.Println(repeatedNTimes([]int{5,1,5,2,5,3,5,4}))
}
