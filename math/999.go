package main

import "fmt"

func numRookCaptures(board [][]byte) int {
    if len(board) == 0 || len(board) != len(board[0]) {
	return 0
    }

    k, t := findRookPosition(board)
    if k < 0 {
	return 0
    }
    return move(board, k, t, 'N') + move(board, k, t, 'S') + move(board, k, t, 'L') + move(board, k, t, 'R')
}

func findRookPosition(board [][]byte) (int, int) {
    for i := 0; i < len(board); i++ {
	for j := 0; j < len(board[i]); j++ {
	    if board[i][j] == 'R' {
		return i, j
	    }
	}
    }

    return -1, -1
}

func move(board [][]byte, k, t int, direction byte) int {
    switch direction {
    case 'N':
	i := k - 1
	for i >= 0 {
	    count := checkChess(board[i][t])
	    if count >= 0 {
		return count
	    }
	    i--
	}
	return 0
    case 'S':
	i := k + 1
	for i < len(board) {
	    count := checkChess(board[i][t])
	    if count >= 0 {
		return count
	    }
	    i++
	}
	return 0
    case 'L':
	j := t - 1
	for j >= 0 {
	    count := checkChess(board[k][j])
	    if count >= 0 {
		return count
	    }
	    j--
	}
	return 0
    case 'R':
	j := t + 1
	for j < len(board[k]) {
	    count := checkChess(board[k][j])
	    if count > 0 {
		return count
	    }
	    j++
	}

	return 0
    default:
	return 0
    }
}

func checkChess(chess byte) int {
    if chess == 'p' {
	return 1
     }

    if chess == 'B' {
        return 0
    }

    return -1
}

func main() {
}
