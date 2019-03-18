package main

import "fmt"

func divide(dividend int, divisor int) int {
    dividendSign,divisorSign := 1, 1
    if dividend < 0 {
	dividend = -dividend
	dividendSign = -1
    }

    if divisor < 0 {
	divisor = - divisor
	divisorSign = -1
    }

    count := 0
    if divisor == 1 {
	result = dividend
    }else {
	for dividend >= divisor {
	    dividend -= divisor
	    count++
	}
    }



    result := dividendSign * divisorSign * count
    if result > ( 1 << 31 - 1) {
	result = 1 << 31 - 1
    }else  if result < -(1 << 31) {
	result = -(1 << 31)
    }

    return result
}

func main() {
    fmt.Println(divide(10, 3))
    fmt.Println(divide(7, -3))
    fmt.Println(divide(1,1))
    fmt.Println(divide(-2147483648,-1))
}
