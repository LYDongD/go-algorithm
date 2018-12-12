//package main
//
//import (
//	"fmt"
//)

func checkPerfectNumber(num int) bool {
	if num < 6 {
		return false
	}

	sum := 0
	i := 2
	j := num / i
	for i <= j {
		if num%i == 0 {
			sum += i
			if j != i {
				sum += j
			}
		}

		i++
		j = num / i
	}

	return num == sum+1
}

func sqrt(num int) int {

	i := 1
	j := num

	for i < j {
		middle := i + (j-i)>>1
		if num/middle == middle {
			return middle
		} else if num/middle > middle {
			i = middle + 1
		} else {
			j = middle - 1
		}
	}

	if num/i > i {
		return i
	} else {
		return num / i
	}
}

//func main() {
//	fmt.Println(checkPerfectNumber(28))
//	fmt.Println(checkPerfectNumber(30))
//	fmt.Println(checkPerfectNumber(6))
//}
