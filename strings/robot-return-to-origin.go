package main

import "fmt"

func judgeCircle(moves string) bool {

    vertical := 0
    horizontal := 0
    for _, char := range moves {
        switch char {
	case 'U':
	    vertical++
	case 'D':
	    vertical--
	case 'L':
	    horizontal++
	case 'R':
	    horizontal--
	}
    }

    return vertical == 0 && horizontal == 0
}


func main() {
    fmt.Println(judgeCircle("UD"))
    fmt.Println(judgeCircle("LL"))
}
