package main

import "fmt"

func divisorGame(N int) bool {
    if N < 2 {
	return false
    }

    return N % 2 == 0
}

func main() {
    fmt.Println(divisorGame(2))
    fmt.Println(divisorGame(3))
    fmt.Println(divisorGame(5))
}
