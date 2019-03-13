package main

import "fmt"

var result [][]string
var board []string

func solveNQueens(n int) [][]string {
    if n == 0 {
	return result
    }
    board = make([]string, n)
    placeQueue(n, 0)
    return result
}

func placeQueue(n int, row int) {

    if row == n {
	lawfulBoard := makeLawfulBoard(board)
	result = append(result, lawfulBoard)
	return
    }

    for i := 0; i < n; i++ {
	if isOk(row, i, n) {
	  board[row] = configuRowStr(i, n)
	  placeQueue(n, row+1)
	}
    }
}

func makeLawfulBoard(board []string) []string{
    lawfulBoard := make([]string,len(board))
    for i := 0; i < len(board); i++ {
	lawfulBoard[i] = board[i]
    }

    return lawfulBoard
}

func configuRowStr(i, n int) string {
    rowStr := ""
    for k := 0; k < n; k++ {
	if k == i {
	    rowStr = rowStr + "Q"
	}else {
	    rowStr = rowStr + "."
	}
    }

    return rowStr
}

func isOk(row, col, n int) bool {
    leftUp := col - 1
    rightUp := col + 1
    for i := row -1; i >= 0; i-- {
	//up
	if board[i][col] == 'Q' {
	    return false
	}

	//left up
	if leftUp >=0 && board[i][leftUp] == 'Q' {
	    return false
	}

	//right up
	if rightUp < n && board[i][rightUp] == 'Q'{
	    return false
	}

	leftUp--
	rightUp++
    }

    return true
}

func main() {
    fmt.Println(solveNQueens(1))
}
