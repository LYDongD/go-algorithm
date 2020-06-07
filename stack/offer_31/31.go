package main

import "fmt"

type Stack struct {
    nums []int
}

func (this *Stack) Push(num int) {
    this.nums = append(this.nums, num)
}

func (this *Stack) Pop() int {
    if len(this.nums) == 0 {
	panic("stack is empty")
    }

    top := this.Top()
    this.nums = this.nums[0:len(this.nums)-1]
    return top
}

func (this *Stack) Top() int {
    return this.nums[len(this.nums) -1]
}

func (this *Stack) IsEmpty() bool {
    return len(this.nums) == 0
}

func validateStackSequences(pushed []int, popped []int) bool {
    if len(pushed) != len(popped) {
	return false
    }

    if len(pushed) == 0 {
	return true
    }

    pushIndex, popIndex := 0, 0
    stack := &Stack{[]int{}}
    for pushIndex < len(pushed) || popIndex < len(popped) {
	if stack.IsEmpty() || stack.Top() != popped[popIndex] {
	    //can't pod and no more ele can be pushe
	    if pushIndex >= len(pushed) {
		return false
	    }
	    stack.Push(pushed[pushIndex])
	    pushIndex++
	}else {
	    stack.Pop()
	    popIndex++
	}
    }

    return true
}

func main() {
    fmt.Println(validateStackSequences([]int{1,2,3,4,5}, []int{4,5,3,2,1}))
    fmt.Println(validateStackSequences([]int{1,2,3,4,5}, []int{4,3,5,1,2}))
}
