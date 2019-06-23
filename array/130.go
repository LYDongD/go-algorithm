package main

import "fmt"

func solve(board [][]byte) {
	if len(solve) == 0 {
		return board
	}

	row, col := len(board), len(board[0])
	notFlippedDic := makeBoardShape(board)
	visited := makeBoardShape(board)

	i, j := 0, 0
	for {
		if !visited[i][j] {
			visited[i][j] = 1
			if board[i][j] == 'o' {
				checkAdjacent(board, visited, notFlippedDic, i, j)
			}
		}

		//iterate boarder clockwise
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
}

func main() {
	fmt.Println("vim-go")
}
