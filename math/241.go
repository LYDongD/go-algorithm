package main

import "fmt"
import "strconv"

func diffWaysToCompute(input string) []int {
    result := []int{}
    if len(input) == 0 {
	return result
    }

    return computeReccusively(input)
}

func computeReccusively(input string) []int {
    result := []int{}
    for i := 0; i < len(input); i++ {
	if !isOperator(input[i]) {
	    continue
	}

	leftArr := computeReccusively(input[0:i])
	rightArr := computeReccusively(input[i+1:len(input)])
	for _, left := range leftArr {
	    for _, right := range rightArr {
		switch input[i] {
		case '+':
		    result = append(result, left + right)
		case '-':
		    result = append(result, left - right)
		case '*':
		    result = append(result, left * right)
		}
	    }
	}
    }

    if len(result) == 0 {
	num, err := strconv.Atoi(input)
	if err == nil {
	    result = append(result, num)
	}
    }

    return result
}

func isOperator(char byte) bool {
    return char == '+' || char == '-' || char == '*'
}

func main() {
    fmt.Println(diffWaysToCompute("2-1-1"))
    fmt.Println(diffWaysToCompute("2*3-4*5"))
}
