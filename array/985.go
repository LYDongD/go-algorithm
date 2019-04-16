package main

import "fmt"

func sumEvenAfterQueries(A []int, queries [][]int) []int {
	if len(A) == 0 || len(queries) == 0 {
		return A
	}

	evenSum := 0
	for _, num := range A {
		if num%2 == 0 {
			evenSum += num
		}
	}

	result := []int{}
	for i := 0; i < len(queries); i++ {
		val := queries[i][0]
		index := queries[i][1]

		if A[index]%2 == 0 {
			evenSum -= A[index]
		}

		A[index] += val
		if A[index]%2 == 0 {
			evenSum += A[index]
		}

		result = append(result, evenSum)
	}

	return result
}

func main() {
	fmt.Println(sumEvenAfterQueries([]int{1, 2, 3, 4}, [][]int{{1, 0}, {-3, 1}, {-4, 0}, {2, 3}}))
}
