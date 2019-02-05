package main

import "fmt"

type MyLinkedList struct {
	head *Node
	tail *Node
	size int
}

//doubly linked list
type Node struct {
	val  int
	next *Node
	prev *Node
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{nil, nil, 0}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	node := this.getNode(index)
	if node != nil {
		return node.val
	}

	return -1
}

func (this *MyLinkedList) getNode(index int) *Node {
	if index <= this.size/2 {
		cursor := this.head
		idx := 0
		for cursor != nil {
			if idx == index {
				return cursor
			}

			idx++
			cursor = cursor.next
		}
		return nil
	}

	if index > this.size/2 {
		cursor := this.tail
		idx := this.size - 1
		for cursor != nil {
			if idx == index {
				return cursor
			}

			idx--
			cursor = cursor.prev
		}

		return nil
	}

	return nil
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {

	cursor := this.head
	if cursor == nil {
		this.head = &Node{val, nil, nil}
		this.tail = this.head
	} else {
		node := &Node{val, cursor, nil}
		this.head = node
		cursor.prev = this.head
	}

	this.size++
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	if this.tail == nil {
		this.tail = &Node{val, nil, nil}
		this.head = this.tail
	} else {
		node := &Node{val, nil, this.tail}
		this.tail.next = node
		this.tail = node
	}

	this.size++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index == 0 {
		this.AddAtHead(val)
		return
	}

	if index == this.size {
		this.AddAtTail(val)
		return
	}

	if index > this.size {
		return
	}

	cursor := this.getNode(index)
	if cursor != nil {
		node := &Node{val, cursor, cursor.prev}
		cursor.prev.next = node
		cursor.prev = node
		this.size++
	}

}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {

	if this.size == 0 {
		return
	}

	if index > this.size-1 {
		return
	}

	if this.size == 1 && index == 0 {
		this.head = nil
		this.tail = nil
		this.size--
		return
	}

	if index == 0 {
		cursor := this.head
		this.head = cursor.next
		this.head.prev = nil
		cursor = nil
		this.size--
		return
	}

	if index == this.size-1 {
		cursor := this.tail
		this.tail = cursor.prev
		this.tail.next = nil
		cursor = nil
		this.size--
		return
	}

	cursor := this.getNode(index)
	cursor.prev.next = cursor.next
	cursor.next.prev = cursor.prev
	cursor = nil
	this.size--
}

func (this *MyLinkedList) printLink() {

	cursor := this.head
	for cursor.next != nil {
		fmt.Print(cursor.val)
		fmt.Print("->")
		cursor = cursor.next
	}

	fmt.Print(cursor.val)
	fmt.Println(" size: ", this.size)
}

func (this *MyLinkedList) printLink2() {
	cursor := this.tail
	for cursor.prev != nil {
		fmt.Print(cursor.val)
		fmt.Print("<-")
		cursor = cursor.prev
	}

	fmt.Print(cursor.val)
	fmt.Println("")
	fmt.Println(" size: ", this.size)
}

func main() {
	obj := Constructor()
	obj.AddAtHead(38)
	obj.AddAtHead(45)
	obj.printLink()

	obj.DeleteAtIndex(2)
	obj.printLink()

	obj.AddAtIndex(1, 24)
	obj.printLink()
}
