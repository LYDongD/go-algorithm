package main

import "fmt"

func lemonadeChange(bills []int) bool {
    if len(bills) == 0 {
	return true
    }

    fiveCount, tenCount, twentyCount := 0, 0, 0
    for _, bill := range bills {
	switch bill {
	case 5:
	    fiveCount++
	case 10:
	    tenCount++
	    fiveCount--
	    if fiveCount < 0 {
		return false
	    }
	case 20:
	    twentyCount++
	    if tenCount > 0 && fiveCount > 0 {
		tenCount--
		fiveCount--
	    }else if fiveCount >= 3 {
		fiveCount = fiveCount - 3
	    }else {
		return false
	    }
	}
    }

    return true
}


func main() {
    fmt.Println(lemonadeChange([]int{5,5,5,10,20}))
    fmt.Println(lemonadeChange([]int{5,5,10}))
    fmt.Println(lemonadeChange([]int{5,5,10,10,20}))
    fmt.Println(lemonadeChange([]int{10,10}))
}
