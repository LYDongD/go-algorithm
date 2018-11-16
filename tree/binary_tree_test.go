package tree

import "testing"

func Test_InorderTraverse(t *testing.T) {
	binaryTree := NewBinaryTree(NewNode(1))
	binaryTree.root.left = NewNode(3)
	binaryTree.root.right = NewNode(4)
	binaryTree.root.right.left = NewNode(5)

	binaryTree.InOrderTraverse()
}

func Test_PreOrderTraverse(t *testing.T) {
	binaryTree := NewBinaryTree(NewNode(1))
	binaryTree.root.left = NewNode(3)
	binaryTree.root.right = NewNode(4)
	binaryTree.root.right.left = NewNode(5)

	binaryTree.PreOrderTraverse()
}

func Test_PostOrderTraverse(t *testing.T) {
	binaryTree := NewBinaryTree(NewNode(1))
	binaryTree.root.left = NewNode(3)
	binaryTree.root.right = NewNode(4)
	binaryTree.root.right.left = NewNode(5)

	binaryTree.PostOrderTraverse()
}
