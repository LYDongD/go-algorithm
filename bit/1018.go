package main

import "fmt"

func prefixesDivBy5(A []int) []bool {
    answer := make([]bool, len(A))
    if len(A) == 0 {
	return nil
    }

    n := 0
    for i := 0; i < len(A); i++ {
	n = (n << 1 + A[i]) % 5
	answer[i] = (n == 0)
    }

    return answer
}


func main() {
    fmt.Println(prefixesDivBy5([]int{0,1,1}))
    fmt.Println(prefixesDivBy5([]int{1,1,1}))
    fmt.Println(prefixesDivBy5([]int{0,1,1,1,1,1}))
}
