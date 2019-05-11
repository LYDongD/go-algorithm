package main

import "fmt"

func exist(board [][]byte, word string) bool {
    if len(board) == 0 || len(board[0]) == 0 || len(word) == 0 {
	return false
    }

    moved := make([][]bool, len(board))
    for r := 0; r < len(moved); r++ {
	moved[r] = make([]bool, len(board[r]))
    }

    //find first char to move 
    result := false
    for i := 0; i < len(board); i++ {
	for j := 0; j < len(board[i]); j++ {
	    if board[i][j] == word[0] {
		moved[i][j] = true
		result  = result || moveRecursively(i, j, 1, moved, board, word)
		moved[i][j] = false
	    }
	}
    }

    return result
}

func moveRecursively(i, j, cursor int, moved [][]bool, board [][]byte, word string) bool {
    if cursor ==  len(word) {
	return true
    }

    //move by 4 directiaon
    result := false
    nextChar := word[cursor]
    canMoveUp := i - 1 >= 0 && board[i-1][j] == nextChar && moved[i-1][j] == false
    if canMoveUp {
	moved[i-1][j] = true
	result = result || moveRecursively(i-1, j, cursor+1, moved, board, word)
	moved[i-1][j] = false
    }

    canMoveDown := i + 1 < len(board) && board[i+1][j] == nextChar && moved[i+1][j] == false
    if canMoveDown {
	moved[i+1][j] = true
	result = result || moveRecursively(i+1, j, cursor+1, moved, board, word)
	moved[i+1][j] = false
    }

    canMoveLeft := j - 1 >= 0  && board[i][j-1] == nextChar && moved[i][j-1] == false
    if canMoveLeft {
	moved[i][j-1] = true
	result = result || moveRecursively(i, j-1, cursor+1, moved, board, word)
	moved[i][j-1] = false
    }

    canMoveRight := j + 1 < len(board[0]) && board[i][j+1] == nextChar && moved[i][j+1] == false
    if canMoveRight {
	moved[i][j+1] = true
	result = result || moveRecursively(i, j+1, cursor+1, moved, board, word)
	moved[i][j+1] = false
    }

    return result
}

func main() {
    fmt.Println(exist([][]byte{{'A','B','C','E'},{'S','F','C','S'},{'A','D','E','E'}}, "ABCCED"))
    fmt.Println(exist([][]byte{{'A','B','C','E'},{'S','F','C','S'},{'A','D','E','E'}}, "SEE"))
    fmt.Println(exist([][]byte{{'A','B','C','E'},{'S','F','C','S'},{'A','D','E','E'}}, "ABCB"))
}
