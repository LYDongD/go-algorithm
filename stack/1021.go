package main

import "fmt"

type Stack struct {
	nums []byte
}

func newStack() *Stack {
	stack := &Stack{nums: []byte{}}
	return stack
}

func (this *Stack) isEmpty() bool {
	return len(this.nums) == 0
}

func (this *Stack) push(num byte) {
	this.nums = append(this.nums, num)
}

func (this *Stack) pop() byte {
	if this.isEmpty() {
		panic("stack is empty!!")
	}

	last := this.nums[len(this.nums)-1]
	this.nums = this.nums[:len(this.nums)-1]

	return last
}

func removeOuterParentheses(S string) string {
	if len(S) == 0 {
		return S
	}

	stack := newStack()
	result := ""
	primitiveStart := 0
	for i := 0; i < len(S); i++ {
		if S[i] == '(' {
			stack.push(S[i])
		} else if S[i] == ')' {
			stack.pop()
			//get primitive sub string
			if stack.isEmpty() {
				primitive := S[primitiveStart : i+1]
				primitiveStart = i + 1

				//delete outest parenthese
				result = result + primitive[1:len(primitive)-1]
			}
		}
	}

	return result
}

func main() {
	fmt.Println(removeOuterParentheses("(()())(())"))
	fmt.Println(removeOuterParentheses("(()())(())(()(()))"))
	fmt.Println(removeOuterParentheses("()()"))
}
