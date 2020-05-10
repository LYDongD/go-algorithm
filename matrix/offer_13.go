package main

import "fmt"

func movingCount(m int, n int, k int) int {
	if m <= 0 || n <= 0 || k < 0 {
		return 0
	}

	visited := newVisitedMatrix(m, n)
	return search(0, 0, m, n, k, visited)
}

func search(i, j, m, n, k int, visited [][]bool) int {
	if sum(i, j) > k || visited[i][j] {
		return 0
	}

	count := 1
	visited[i][j] = true

	if i+1 <= m-1 {
		count = count + search(i+1, j, m, n, k, visited)
	}

	if i-1 >= 0 {
		count = count + search(i-1, j, m, n, k, visited)
	}

	if j+1 <= n-1 {
		count = count + search(i, j+1, m, n, k, visited)
	}

	if j-1 >= 0 {
		count = count + search(i, j-1, m, n, k, visited)
	}

	return count
}

func sum(a, b int) int {
	s := 0
	for a > 0 {
		s += a % 10
		a = a / 10
	}

	for b > 0 {
		s += b % 10
		b = b / 10
	}

	return s
}

func newVisitedMatrix(m, n int) [][]bool {
	result := make([][]bool, m)
	for i := 0; i < m; i++ {
		result[i] = make([]bool, n)
	}
	return result
}

func main() {
	fmt.Println(movingCount(3, 1, 0))
	fmt.Println(movingCount(2, 3, 1))
	fmt.Println(movingCount(3, 2, 17))
}
