package main

import "fmt"

func singleNumber(nums []int) []int {
    result := []int{}
    if len(nums) < 2 {
	return result
    }

    diff := 0
    for _, num := range nums  {
	diff = diff ^ num
    }

    //find the right-most diff position
    mask := diff & (-diff)

    //depart single num by &
    single1,single2 := 0,0
    for _, num := range nums {
	if mask & num == 0 {
	    single1 = single1 ^ num
	}else {
	    single2 = single2 ^ num
	}
    }

    result = append(result, single1)
    result = append(result, single2)

    return result
}

func main() {
    fmt.Println(singleNumber([]int{1,2,1,3,2,5}))
}
