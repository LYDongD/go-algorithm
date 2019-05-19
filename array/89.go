package main

import "fmt"

func grayCode(n int) []int {

    result := []int{0}
    if n == 0 {
	return result
    }

    //generate gray code recursively
    for i := 0; i < n; i++ {
	for k := len(result) - 1; k >= 0; k-- {
	    result = append(result, result[k] | (1 << uint(i)))
	}
    }

    return result
}


func main() {
    fmt.Println(grayCode(3))
    fmt.Println(grayCode(2))
}
