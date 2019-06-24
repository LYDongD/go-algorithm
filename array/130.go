package main

import "fmt"

func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}

	row, col := len(board), len(board[0])
	notFlippedDic := makeBoardShape(board)
	visited := makeBoardShape(board)

	i, j := 0, 0
	for {
		if board[i][j] == 'O' {
			notFlippedDic[i][j] = 1
			checkAdjacent(board, visited, notFlippedDic, i, j)
		}

		//iterate board clockwise
		if i == 0 && j < col-1 {
			j++
		} else if j == col-1 && i < row-1 {
			i++
		} else if i == row-1 && j > 0 {
			j--
		} else if j == 0 && i > 0 {
			i--
		}

		//back to start point
		if j == 0 && i == 0 {
			break
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if notFlippedDic[i][j] == 0 {
				board[i][j] = 'X'
			}
		}
	}
}

func makeBoardShape(board [][]byte) [][]int {
	result := make([][]int, len(board))
	for i := 0; i < len(result); i++ {
		result[i] = make([]int, len(board[i]))
	}

	return result
}

func checkAdjacent(board [][]byte, visited, notFlippedDic [][]int, i, j int) {

	if visited[i][j] == 1 {
		return
	}

	visited[i][j] = 1
	if board[i][j] == 'O' {
		notFlippedDic[i][j] = 1

		//check around
		if j > 0 {
			checkAdjacent(board, visited, notFlippedDic, i, j-1)
		}

		if j < len(board[0])-1 {
			checkAdjacent(board, visited, notFlippedDic, i, j+1)
		}

		if i > 0 {
			checkAdjacent(board, visited, notFlippedDic, i-1, j)
		}

		if i < len(board)-1 {
			checkAdjacent(board, visited, notFlippedDic, i+1, j)
		}
	}
}

func main() {
	//source := [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}}
	source := [][]byte{{'O', 'O', 'O'}, {'O', 'O', 'O'}, {'O', 'O', 'O'}}
	solve(source)
	fmt.Println(source)
}
