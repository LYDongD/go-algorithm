package main

import (
	"fmt"
	"strconv"
)

func findRelativeRanks(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}

	//sort
	sortedNums := sortNums(nums)

	rankedMap := generateRankedMap(sortedNums)

	result := make([]string, len(nums))
	for index, num := range nums {
		rank, ok := rankedMap[num]
		if ok {
			switch rank {
			case 1:
				result[index] = "Gold Medal"
			case 2:
				result[index] = "Silver Medal"
			case 3:
				result[index] = "Bronze Medal"
			default:
				result[index] = strconv.Itoa(rank)
			}
		}
	}

	return result

}

//bubble sort
func sortNums(nums []int) []int {
	a := make([]int, len(nums))
	copy(a, nums)

	quickSort(a, 0, len(a)-1)
	return a
}

func quickSort(a []int, p int, r int) {

	if p < r {
		//use i to mark the left side last index
		q := partion(a, p, r)
		quickSort(a, p, q-1)
		quickSort(a, q, r)
	}
}

func partion(a []int, p int, r int) int {

	pivort := a[r]
	i := p - 1
	//move the bigger num to leftside
	for j := p; j < r; j++ {
		if a[j] > pivort {
			i++
			//swap
			a[j], a[i] = a[i], a[j]
		}
	}

	//swap pivort and the middle
	a[i+1], a[r] = a[r], a[i+1]

	//return the middle
	return i + 1
}

func generateRankedMap(sortedNums []int) map[int]int {
	rankedMap := make(map[int]int)
	for index, num := range sortedNums {
		rankedMap[num] = index + 1
	}
	return rankedMap
}

func main() {

	nums := []int{8, 1, 5, 2, 3}
	//sortedNums := sortNums(nums)
	//fmt.Println(sortedNums)
	//fmt.Println(generateRankedMap(sortedNums))
	fmt.Println(findRelativeRanks(nums))
}
