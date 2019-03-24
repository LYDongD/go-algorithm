package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
    if len(board) == 0 || len(board) != len(board[0]) {
	return false
    }

    //row check
    for i := 0; i < len(board); i++ {
	rowDic := make(map[byte]int)
	for j := 0; j < len(board[i]); j++{
	    if board[i][j] != '.' && rowDic[board[i][j]] > 0 {
		return false
	    }
	    rowDic[board[i][j]]++
	}

    }

    //col check
    for j := 0; j < len(board[0]); j++ {
	colDic := make(map[byte]int)
	for i := 0; i < len(board); i++ {
	    if board[i][j] != '.' && colDic[board[i][j]] > 0 {
		return false
	    }

	    colDic[board[i][j]]++
	}
    }

    //sub box check
    for i := 0; i < len(board); i = i + 3 {
	for j := 0; j < len(board[i]); j = j + 3 {
	    boxDic := make(map[byte]int)
	    for k := i; k < i + 3; k++ {
		for t := j; t < j + 3; t++ {
		    if board[k][t] != '.' && boxDic[board[k][t]] > 0 {
			return false
		    }
		    boxDic[board[k][t]]++
		}
	    }
	}
    }

    return true
}

func main() {
    fmt.Println("ss")
}

