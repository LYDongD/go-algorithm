package main

import "fmt"

func minDeletionSize(A []string) int {
    if len(A) <= 1 {
	return 0
    }

    colums := len(A[0])
    count := 0
    for i := 0; i < colums; i++ {
	for j := 1; j < len(A); j++ {
	    if A[j][i] < A[j-1][i] {
		count++
		break
	    }
	}
    }

    return count
}

func main() {
    fmt.Println(minDeletionSize([]string{"cba","daf","ghi"}))
    fmt.Println(minDeletionSize([]string{"a","b"}))
    fmt.Println(minDeletionSize([]string{"zyx","wvu","tsr"}))
}
