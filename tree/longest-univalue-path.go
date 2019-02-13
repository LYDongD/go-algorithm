package main


func longestUnivaluePath(root *TreeNode) int {
    max := 0
    findPath(root, &max)
    return max
}

func findPath(node *TreeNode, max *int) int{
    if node == nil {
        return 0
    }

    left := findPath(node.Left, max)
    right := findPath(node.Right, max)
    leftPath := 0
    rightPath := 0

    if node.Left != nil && node.Val == node.Left.Val{
        leftPath = left + 1
    }

    if node.Right != nil && node.Val == node.Right.Val{
        rightPath =  right + 1
    }

    if *max < leftPath + rightPath {
        *max = leftPath + rightPath
    }
    return maxValue(leftPath, rightPath)
}

func maxValue(a int, b int) int{
    if a >= b {
        return a
    }else {
        return b
    }
}


