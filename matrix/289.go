package main

import "fmt"

func gameOfLife(board [][]int)  {
    if len(board) == 0 || len(board[0]) == 0 {
	return
    }

    for i := 0; i < len(board); i++ {
	for j := 0; j < len(board[i]); j++ {
	    if board[i][j] == 0 && canLive(board, i, j){
		board[i][j] = -1
		continue
	    }

	    if board[i][j] == 1 && willDie(board, i, j) {
		board[i][j] = 2
	    }
	}
    }

    for i := 0; i < len(board); i++ {
	for j := 0; j < len(board[i]); j++ {
	    if board[i][j] == 2 {
		board[i][j] = 0
	    }else if board[i][j] == -1 {
		board[i][j] = 1
	    }
	}
    }
}

func canLive(board [][]int, i, j int) bool {
    live, _ := countLiveAndDead(board, i, j)
    return live == 3
}

func willDie(board [][]int, i, j int) bool {
    live, _ := countLiveAndDead(board, i, j)
    return live > 3 || live < 2
}

func countLiveAndDead(board [][]int, i, j int) (int, int)  {

    liveCount := 0
    deadCount := 0

    //up
    if i - 1 >= 0 {
	if board[i-1][j] > 0 {
	    liveCount++
	}else {
	    deadCount++
	}
    }

    //down
    if i + 1 < len(board) {
	if board[i+1][j] > 0 {
	    liveCount++
	}else {
	    deadCount++
	}
    }

    //LEFT
    if j - 1 >= 0 {
	if board[i][j-1] > 0 {
	    liveCount++
	}else {
	    deadCount++
	}
    }

    //right
    if j + 1 < len(board[0]) {
	if board[i][j+1] > 0 {
	    liveCount++
	}else {
	    deadCount++
	}
    }


    //left up
    if i - 1 >= 0 && j - 1 >= 0 {
	if board[i-1][j-1] > 0 {
	    liveCount++
	}else {
	    deadCount++
	}
    }

    //left bottom
    if i + 1 < len(board)  && j - 1 >= 0 {
	if board[i+1][j-1] > 0 {
	    liveCount++
	}else {
	    deadCount++
	}
    }
    //right up
    if i - 1 >= 0 && j + 1 < len(board[0])  {
	if board[i-1][j+1] > 0 {
	    liveCount++
	}else {
	    deadCount++
	}
    }
    //right bottom
    if i + 1 < len(board)  && j + 1 < len(board[0])  {
	if board[i+1][j+1] > 0 {
	    liveCount++
	}else {
	    deadCount++
	}
    }

    return liveCount, deadCount
}



func main() {
    board := [][]int{{0,1,0},{0,0,1},{1,1,1},{0,0,0}}
    gameOfLife(board)
    fmt.Println(board)

}
