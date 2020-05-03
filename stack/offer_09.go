package main

import "fmt"

type CQueue struct {
	stack1 *Stack
	stack2 *Stack
}

type Stack struct {
	nums []int
}

func (self *Stack) Push(num int) {
	self.nums = append(self.nums, num)
}

func (self *Stack) Pop() int {
	if self.IsEmpty() {
		panic("stack is empty")
	}

	top := self.nums[len(self.nums)-1]
	self.nums = self.nums[0 : len(self.nums)-1]
	return top
}

func (self *Stack) IsEmpty() bool {
	return len(self.nums) == 0
}

func Constructor() CQueue {
	queue := CQueue{stack1: &Stack{[]int{}}, stack2: &Stack{[]int{}}}
	return queue
}

func (this *CQueue) AppendTail(value int) {
	for !this.stack1.IsEmpty() {
		this.stack2.Push(this.stack1.Pop())
	}

	this.stack1.Push(value)

	for !this.stack2.IsEmpty() {
		this.stack1.Push(this.stack2.Pop())
	}
}

func (this *CQueue) DeleteHead() int {
	if this.stack1.IsEmpty() {
		return -1
	}

	return this.stack1.Pop()
}

func main() {
	fmt.Println("running")
	queue := Constructor()
	queue.AppendTail(3)
	queue.AppendTail(5)
	queue.AppendTail(7)
	fmt.Println(queue.DeleteHead())
	fmt.Println(queue.DeleteHead())
	fmt.Println(queue.DeleteHead())
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
