package main

import "fmt"

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
    top := self.nums[len(self.nums)-1]
    self.nums = self.nums[0:len(self.nums)-1]
    return top
}

func calculate(s string) int {
    if len(s) == 0 {
	return 0
    }

    //延迟入栈，当前数终止时，根据前置符号决定入栈数
    num := 0
    sign := '+'
    stack := &Stack{[]int{}}
    for i, char := range s {

	//lazy push
	if isDigit(char) {
	    num = num * 10 + int(char - '0')
	}

	//push by sign
	if isOperator(char) || i == len(s) -1 {
	    if sign == '+' {
		stack.Push(num)
	    }else if sign == '-' {
		stack.Push(-num)
	    }else if sign == '*' {
		stack.Push(stack.Pop() * num)
	    }else if sign == '/' {
		stack.Push(stack.Pop() / num)
	    }
	    num = 0
	    sign = char
	}
    }

    result := 0
    for !stack.IsEmpty() {
	result += stack.Pop()
    }

    return result

}

func isDigit(char rune) bool {
    return char >= '0' && char <= '9'
}

func isOperator(char rune) bool {
    return char == '+' || char == '-' || char == '*' || char == '/'
}

func main() {
    fmt.Println(calculate("3+2*2"))
    fmt.Println(calculate(" 3/2 "))
    fmt.Println(calculate(" 3+5 / 2 "))
}
