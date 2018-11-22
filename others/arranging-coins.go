//package main

//import "fmt"

func arrangeCoins(n int) int {
	sum := 0
	rowCoinCount := 0
	for {
		if rowCoinCount+1 <= n-sum {
			rowCoinCount++
			sum += rowCoinCount
		} else {
			return rowCoinCount
		}
	}
}

//func main() {
//	fmt.Println(arrangeCoins(5))
//	fmt.Println(arrangeCoins(8))
//}
