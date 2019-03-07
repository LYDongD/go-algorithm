package main

import "fmt"

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
    if len(rec1) < 4 || len(rec2) < 4 {
	return false
    }

    x1,y1,x2,y2 := rec1[0], rec1[1], rec1[2], rec1[3]
    a1,b1,a2,b2 := rec2[0], rec2[1], rec2[2], rec2[3]

    if x1 >= a2 || x2 <= a1 || y1 >= b2 || y2 <= b1 {
	return false
    }

    return true
}

func main() {
    fmt.Println(isRectangleOverlap([]int{0,0,2,2}, []int{1,1,3,3}))
    fmt.Println(isRectangleOverlap([]int{0,0,1,1}, []int{1,0,2,1}))
}
