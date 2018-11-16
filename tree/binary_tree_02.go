package tree

import (
	"fmt"
)

type BinaryTree struct {
	root *Node
}

func NewBinaryTree(root *Node) *BinaryTree {
	return &BinaryTree{root: root}
}

//traverse by stack
func (this *BinaryTree) InOrderTraverse() {
	r := this.root
	s := NewArrayStack()

	for !s.IsEmpty() || r != nil {
		if r != nil {
			s.Push(r)
			r = r.left
		} else {
			tmp := s.Pop().(*Node)
			fmt.Println(tmp.data)
			r = tmp.right
		}
	}
}

func (this *BinaryTree) PreOrderTraverse() {
	r := this.root
	s := NewArrayStack()

	for !s.IsEmpty() || r != nil {
		if r != nil {
			fmt.Println(r.data)
			s.Push(r)
			r = r.left
		} else {
			r = s.Pop().(*Node).right
		}

	}
}

func (this *BinaryTree) PostOrderTraverse() {
	r := this.root
	s := NewArrayStack()

	//point to last visit node
	var pre *Node

	s.Push(r)

	for !s.IsEmpty() {
		r = s.Top().(*Node)
		if (r.left == nil && r.right == nil) ||
			(pre != nil && (pre == r.left || pre == r.right)) {

			fmt.Println(r.data)
			s.Pop()
			pre = r
		} else {
			if r.right != nil {
				s.Push(r.right)
			}

			if r.left != nil {
				s.Push(r.left)
			}

		}

	}

}
