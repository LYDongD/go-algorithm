package main

import "fmt"

func isLongPressedName(name string, typed string) bool {
    if len(name) == 0 || len(typed) == 0 {
	return false
    }

    i, j := 0, 0
    for i < len(name) && j < len(typed) {
	if name[i] == typed[j] {
	    i++
	    j++
	}else {
	    if j > 0 && typed[j] == typed[j-1] {
		j++
	    }else {
		return false
	    }
	}
    }

    if j == len(typed) && i < len(name) {
	return false
    }

    if i == len(name) && j < len(typed) {
	j++
	for j < len(typed) {
	    if typed[j] != typed[j-1] {
		return false
	    }
	}
    }

    return true
}

func main() {
    fmt.Println(isLongPressedName("alex", "aaleex"))
    fmt.Println(isLongPressedName("saeed", "ssaaedd"))
    fmt.Println(isLongPressedName("leelee", "lleeelee"))
    fmt.Println(isLongPressedName("laiden", "laiden"))
}
