package main

import "fmt"
import "strings"
import "strconv"

type Stack struct {
	elements []string
}

func (self *Stack) IsEmpty() bool {
	return len(self.elements) == 0
}

func (self *Stack) Push(ele string) {
	self.elements = append(self.elements, ele)
}

func (self *Stack) Pop() string {
	top := self.elements[len(self.elements)-1]
	self.elements = self.elements[0 : len(self.elements)-1]
	return top
}

func (self *Stack) reverse() *Stack {
	newStack := &Stack{[]string{}}
	for !self.IsEmpty() {
		newStack.Push(self.Pop())
	}

	return newStack
}

func calculate(s string) int {
	if len(s) == 0 {
		return 0
	}

	s = strings.Replace(s, " ", "", -1)
	if len(s) == 0 {
		return 0
	}

	stack := &Stack{[]string{}}
	i := 0
	for i <= len(s)-1 {
		if isMultiOperator(s[i : i+1]) {
			num, next := getNextNum(s, i+1)
			result := doOperator(stack.Pop(), s[i:i+1], num)
			i = next
			stack.Push(result)
		} else if isAddOperator(s[i : i+1]) {
			stack.Push((s[i : i+1]))
			i++
		} else {
			num, next := getNextNum(s, i)
			stack.Push(num)
			i = next
		}
	}

	//倒转栈
	stack = stack.reverse()

	for !stack.IsEmpty() {
		a := stack.Pop()
		if stack.IsEmpty() {
			result, err := strconv.Atoi(a)
			if err == nil {
				return result
			}
		}

		operator := stack.Pop()
		b := stack.Pop()
		stack.Push(doOperator(a, operator, b))
	}

	return 0

}

func isMultiOperator(char string) bool {
	return char == "*" || char == "/"
}

func isAddOperator(char string) bool {
	return char == "+" || char == "-"
}

func getNextNum(s string, start int) (string, int) {
	num := ""
	i := 0
	for i = start; i < len(s); i++ {
		if isMultiOperator(s[i:i+1]) || isAddOperator(s[i:i+1]) {
			break
		}
		num = num + s[i:i+1]
	}

	return num, i
}

func doOperator(a, operator, b string) string {
	result := 0
	aint, err := strconv.Atoi(a)
	if err != nil {
		panic("a is not a num")
	}
	bint, err := strconv.Atoi(b)
	if err != nil {
		panic("b is not a num")
	}
	if operator == "+" {
		result = aint + bint
	} else if operator == "-" {
		result = aint - bint
	} else if operator == "*" {
		result = aint * bint
	} else {
		result = aint / bint
	}

	return strconv.Itoa(result)
}

func main() {
	fmt.Println(calculate("42"))
	fmt.Println(calculate("1-1+1"))
	fmt.Println(calculate("3+2*2"))
	fmt.Println(calculate(" 3/2 "))
	fmt.Println(calculate(" 3+5 / 2 "))
	fmt.Println(calculate("0-2147483647"))

}
