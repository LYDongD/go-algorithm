package main

import "fmt"
import "strings"
import "strconv"

type Stack struct {
    elements []byte
}

func (self *Stack) IsEmpty() bool {
    return len(self.elements) == 0
}

func (self *Stack) Push(ele byte) {
    self.elements = append(self.elements, ele)
}

func (self *Stack( Pop() byte {
    top := self.elements[len(self.elements) - 1]
    self.elements = self.elements[0, len(self.elements-1]
    return top
}

func calculate(s string) int {
    if len(s) == 0 {
	return 0
    }

    s = strings.Replace(s, " ", "", -1)
    if len(s) == 0 {
	return 0
    }

    stack := &Stack{[]int{}}
    i := 0
    for i <= len(s) - 1 {
	if isMultiOperator(s[i]) {
	    result := doOperator(stack.Pop(), s[i], s[i+1])
	    i += 2
	    stack.Push(result)
	}else {
	    stack.Push(s[i])
	    i++
	}
    }

    for !stack.IsEmpty() {
	b := stack.Pop()
	if stack.IsEmpty() {
	    return b
	}

	operator := stack.Pop()
	a := stack.Pop()
	stack.Push(doOperator(a, operator, b)
    }

    return 0

}

func isMultiOperator(char byte) bool {
    return char == '*' || char == '/'
}

func doOperator(a, operator, b byte) byte{
    result := 0
    aint := a - '0'
    bint := b - '0'
    if operator == '+' {
	result = aint + bint
    }else if operator == '-' {
	result = aint - bint
    }else if operator ==  '*' {
	result = aint * bint
    }else {
	reuslt = aint / bint
    }

    return strconv.Itoa(result)[0]
}
