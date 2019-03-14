package main

import "fmt"

func generateParenthesis(n int) []string {
    var result []string
    if n <= 0 {
	return result
    }

    return backTrace(n, 0, 0, "", result)
}

func backTrace(n, leftCount, rightCount int, combineStr string, result []string) []string {
    if leftCount == rightCount && leftCount == n {
	result = append(result, combineStr)
	return result
    }

    if leftCount == n {
	result = backTrace(n, leftCount, rightCount + 1, combineStr + ")", result)
    }else if rightCount >= leftCount {
	result = backTrace(n, leftCount + 1, rightCount, combineStr + "(", result)
    }else {
	result = backTrace(n, leftCount + 1, rightCount, combineStr + "(", result)
	result = backTrace(n, leftCount, rightCount + 1, combineStr + ")", result)
    }

    return result
}

func main() {
    fmt.Println(generateParenthesis(4))
    fmt.Println(generateParenthesis(3))
}
