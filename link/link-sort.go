package main

import "fmt"

type Node struct {
	val  int
	prev *Node
	next *Node
}

func sort(head *Node) *Node {
	cursor := head

	for cursor.next != nil {
		cursor = cursor.next
	}

	tail := cursor
	cursor = head

	for tail != nil {
		for cursor.next != nil {
			if cursor.next.val < cursor.val {
				currentVal := cursor.val
				cursor.val = cursor.next.val
				cursor.next.val = currentVal
			}

			cursor = cursor.next
		}
		tail = tail.prev
		cursor = head
	}

	return head
}

func main() {
	list := configureLinkedList([]int{5, 2, 7, 3, 6})
	printList(list)
	fmt.Println("----")
	printList(sort(list))
}

func configureLinkedList(nums []int) *Node {

	var head *Node
	var cursor *Node
	for _, num := range nums {
		node := &Node{num, nil, nil}
		if cursor == nil {
			cursor = node
			head = node
		} else {
			cursor.next = node
			node.prev = cursor
			cursor = cursor.next
		}
	}

	return head
}

func printList(head *Node) {
	cursor := head
	for cursor != nil {
		fmt.Print(cursor.val, ",")
		cursor = cursor.next
	}
	fmt.Println("")
}
