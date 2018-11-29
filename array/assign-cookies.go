package main

import "fmt"

func findContentChildren(g []int, s []int) int {

	if len(g) == 0 || len(s) == 0 {
		return 0
	}

	count := 0

	for _, greed := range g {
		min, minIndex := findMinSizeCookie(greed, s)
		fmt.Println("greed: ", greed, " minIndex: ", minIndex, " min: ", min)
		if minIndex >= 0 {
			count++
			s[minIndex] = -1
		}
	}

	return count
}

func findMinSizeCookie(g int, s []int) (int, int) {
	min := 0
	minIndex := -1
	for index, size := range s {
		if size >= g {
			if min == 0 || min > size {
				min = size
				minIndex = index
			}
		}

	}

	return min, minIndex
}

func main() {

	s := []int{1, 1}
	g := []int{1, 2, 3}

	fmt.Println(findContentChildren(g, s))
}
