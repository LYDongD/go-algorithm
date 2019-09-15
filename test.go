package main

import "fmt"

func main() {
	fmt.Println("hello world")
	for range []int{1, 2, 3} {
		fmt.Println("hi")
	}
	fmt.Println("hello world")
}
