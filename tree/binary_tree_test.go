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

func Test_BinaryTreeBasedArray(t *testing.T) {

	source := []int{2, 6, 3, 7, 9}
	binaryTreeArray := CreateBinaryTree(source)
	for i := range binaryTreeArray {
		t.Log(binaryTreeArray[i])
	}

}
