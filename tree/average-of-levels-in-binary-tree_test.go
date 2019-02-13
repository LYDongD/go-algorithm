package tree

import "testing"

//func Test_queue(t *testing.T) {
//    queue := NewListQueue()
//    t.Log(queue.IsEmpty())
//    queue.Offer(1)
//    queue.Offer(2)
//    t.Log(queue.IsEmpty())
//    t.Log(queue.Poll())
//    t.Log(queue.Poll())
//    t.Log(queue.IsEmpty())
//} 

func Test_averageOfLevels(t *testing.T) {
    root := &TreeNode{Val:3}
    root.Left = &TreeNode{Val:9}
    root.Right = &TreeNode{Val:20}
    root.Right.Left = &TreeNode{Val:15}
    root.Right.Right = &TreeNode{Val:7}

    result := AverageOfLevels(root)
    t.Log(result)
}
