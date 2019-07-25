package main

import "fmt"
import "strconv"

type Stack struct {
    nums []int
}

func (self *Stack) IsEmpty() bool {
    return len(self.nums) == 0
}

func (self *Stack) Push(num int) {
    self.nums = append(self.nums, num)
}

func (self *Stack) Pop() int {
    last := self.nums[len(self.nums)-1]
    self.nums = self.nums[:len(self.nums)-1]
    return last
}


func evalRPN(tokens []string) int {
    if len(tokens) == 0 {
	return 0
    }

    stack := &Stack{nums:[]int{}}
    for _, token := range tokens {
	if !isOperator(token) {
	    tokenInt, err := strconv.Atoi(token)
	    if err == nil {
		stack.Push(tokenInt)
	    }
	}else {
	    right := stack.Pop()
	    left := stack.Pop()
	    stack.Push(doOperation(left, right, token))
	}
    }

    return stack.Pop()
}

func isOperator(token string) bool {
    return token == "+" || token == "-" || token == "*" || token == "/"
}

func doOperation(a, b int, operator string) int {
    if operator == "+" {
	return a + b
    }

    if operator == "-" {
	return a - b
    }

    if operator == "*" {
	return a * b
    }

    if operator == "/" {
	return a / b
    }

    panic("illegal operator")
}


func main() {
    fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))
    fmt.Println(evalRPN([]string{"4", "13", "5", "/", "+"}))
    fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}
