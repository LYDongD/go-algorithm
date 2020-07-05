package main

import "fmt"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	cursor := head
	dummy := &Node{-1, nil, nil}
	cursor_cp := dummy
	node_rel := make(map[*Node]*Node)

	//cp next and mapping one2one
	for cursor != nil {
		cursor_cp.Next = &Node{cursor.Val, nil, nil}
		cursor_cp = cursor_cp.Next
		node_rel[cursor] = cursor_cp
		cursor = cursor.Next
	}

	//cp random
	cursor = head
	cursor_cp = dummy.Next
	for cursor != nil {
		if cursor.Random != nil {
			cursor_cp.Random = node_rel[cursor.Random]
		}
		cursor = cursor.Next
		cursor_cp = cursor_cp.Next
	}

	return dummy.Next

}

func main() {
	fmt.Println("go")
}
