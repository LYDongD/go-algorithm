package main

import "fmt"

func nextGreatestLetter(letters []byte, target byte) byte {
	if len(letters) == 0 {
		panic("no letters found")
	}

	//bounds handle
	if target < letters[0] || target >= letters[len(letters)-1] {
		return letters[0]
	}

	//binary search: 左闭右开区间
	start := 0
	end := len(letters)
	for start < end {
		middle := start + (end-start)/2
		if letters[middle] <= target {
			start = middle + 1
		} else {
			end = middle
		}
	}

	return letters[start]
}

func main() {
	fmt.Println(string(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'a')))
	fmt.Println(string(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'c')))
	fmt.Println(string(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'd')))
	fmt.Println(string(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'g')))
	fmt.Println(string(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'j')))
	fmt.Println(string(nextGreatestLetter([]byte{'c', 'f', 'j'}, 'k')))
}
