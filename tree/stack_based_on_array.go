package tree

import "fmt"

//stack based on array
type ArrayStack struct {
	data []interface{}
	top  int
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		data: make([]interface{}, 0, 32),
		top:  -1,
	}

}

func (this *ArrayStack) IsEmpty() bool {
	if this.top == -1 {
		return true
	}
	return false
}

func (this *ArrayStack) Push(v interface{}) {
	this.top += 1

	//check array compacity
	if this.top > len(this.data)-1 {
		this.data = append(this.data, v)
	} else {
		this.data[this.top] = v
	}
}

func (this *ArrayStack) Pop() interface{} {
	if this.IsEmpty() {
		return nil
	}

	v := this.data[this.top]
	this.top -= 1
	return v
}

func (this *ArrayStack) Top() interface{} {
	if this.IsEmpty() {
		return nil
	}

	return this.data[this.top]

}

func (this *ArrayStack) Flush() {
	this.top = -1
}

func (this *ArrayStack) Print() {
	if this.IsEmpty() {
		fmt.Println("empty stack")
	} else {
		for i := this.top; i >= 0; i++ {
			fmt.Println(this.data[i])
		}
	}
}
