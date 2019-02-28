package main

import "fmt"

func maxArea(height []int) int {
    if len(height) <= 1 {
	return 0
    }

    max := 0
    start := 0
    end := len(height) - 1
    for start < end {
	w := end - start
	h := min(height[end], height[start])
	if w * h > max {
	    max  = w * h
	}

	if height[start] < height[end] {
	    start++
	}else {
	    end--
	}
    }


    return max
}

func min(a, b int) int {
    if a < b {
	return a
    }

    return b
}


func main() {
    fmt.Println(maxArea([]int{1,8,6,2,5,4,8,3,7}))
}
