package tree

func pathSum(root *Node, sum int) int {

	if root == nil {
		return 0
	}
	return nodeForSumCount(root, sum) + pathSum(root.left, sum) + pathSum(root.right, sum)
}

func nodeForSumCount(node *Node, sum int) int {

	if node == nil {
		return 0
	}

	count := 0
	if node.data.(int) == sum {
		count++
	}

	count += nodeForSumCount(node.left, sum-node.data.(int)) + nodeForSumCount(node.right, sum-node.data.(int))

	return count

}
