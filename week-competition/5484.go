package main

import "fmt"

func findKthBit(n int, k int) byte {
	if n < 0 {
		panic("n can't be negative")
	}

	reverse := 0
	for n > 1 {
		l := 1<<uint(n) - 1
		mid := l >> 1
		if k == mid+1 {
			if reverse%2 == 0 {
				return '1'
			} else {
				return '0'
			}
		} else if k > mid+1 {
			k = k - mid - 1
			k = mid - k + 1
			reverse++
		}

		n--
	}

	if reverse%2 == 0 {
		return '0'
	} else {
		return '1'
	}

}
func main() {
	fmt.Println(findKthBit(3, 1))
	fmt.Println(findKthBit(4, 11))
	fmt.Println(findKthBit(1, 1))
	fmt.Println(findKthBit(2, 3))
}
