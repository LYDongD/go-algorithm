package main

import "fmt"

func reorderLogFiles(logs []string) []string {
    if len(logs) == 0 {
	return logs
    }

    result := []string{}
    for _, log := range logs {
	if isDigitLog(log) {
	    result = append(result, log)
	}else {
	    insertedFlag := false
	    for index,loged := range result {
		if isDigitLog(loged) || compare(log, loged) < 0 {
		    right := append([]string{}, result[index:]...)
		    result = append(result[:index], log)
		    result = append(result, right...)
		    insertedFlag = true
		    break
		}
	    }

	    if !insertedFlag {
		result = append(result, log)
	    }
	}
    }

    return result
}

func isDigitLog(log string) bool {
    letterIndex := getFirstLetterIndex(log)
    return log[letterIndex] >= '0' && log[letterIndex] <= '9' 
}

func getFirstLetterIndex(log string) int {
    for index, char := range log {
        if char == ' ' {
            return index + 1
        }
    }

    return -1
}

func compare(a, b string) int {

    letterIndexA := getFirstLetterIndex(a)
    letterIndexB := getFirstLetterIndex(b)
    identifierA, letterA := a[:letterIndexA - 1], a[letterIndexA:]
    identifierB, letterB := b[:letterIndexB - 1], b[letterIndexB:]

    if letterA > letterB {
	return 1
    }

    if letterA < letterB {
	return -1
    }

    if identifierA > identifierB {
	return 1
    }else {
	return -1
    }
}

func main() {
    fmt.Println(reorderLogFiles([]string{"a1 9 2 3 1","g1 act car","zo4 4 7","ab1 off key dog","a8 act zoo"}))
}
