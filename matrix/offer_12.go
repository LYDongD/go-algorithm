package main

import "fmt"

func exist(board [][]byte, word string) bool {
	m, n, k := len(board), len(board[0]), len(word)
	if m == 0 || n == 0 || k == 0 {
		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				visited := newVisitedMatrix(m, n)
				if search(i, j, 0, board, visited, word) {
					return true
				}
			}
		}
	}

	return false
}

func newVisitedMatrix(m, n int) [][]bool {
	result := make([][]bool, m)
	for i := 0; i < m; i++ {
		result[i] = make([]bool, n)
	}
	return result
}

func search(i, j, index int, board [][]byte, visited [][]bool, word string) bool {
	if index == len(word)-1 {
		return true
	}

	visited[i][j] = true
	result := false
	if i+1 <= len(board)-1 && !visited[i+1][j] && board[i+1][j] == word[index+1] {
		result = search(i+1, j, index+1, board, copyMatric(visited), word)
	}
	if !result && i-1 >= 0 && !visited[i-1][j] && board[i-1][j] == word[index+1] {
		result = search(i-1, j, index+1, board, copyMatric(visited), word)
	}
	if !result && j+1 <= len(board[0])-1 && !visited[i][j+1] && board[i][j+1] == word[index+1] {
		result = search(i, j+1, index+1, board, copyMatric(visited), word)
	}
	if !result && j-1 >= 0 && !visited[i][j-1] && board[i][j-1] == word[index+1] {
		result = search(i, j-1, index+1, board, copyMatric(visited), word)
	}

	return result
}

func copyMatric(a [][]bool) [][]bool {
	duplicate := make([][]bool, len(a))
	for i := range a {
		duplicate[i] = make([]bool, len(a[i]))
		copy(duplicate[i], a[i])
	}
	return duplicate
}

func main() {
	board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	fmt.Println(exist(board, "ABCCED"))
	board = [][]byte{{'a', 'b'}, {'c', 'd'}}
	fmt.Println(exist(board, "cdba"))
	board = [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'E', 'S'}, {'A', 'D', 'E', 'E'}}
	fmt.Println(exist(board, "ABCESEEEFS"))
}
