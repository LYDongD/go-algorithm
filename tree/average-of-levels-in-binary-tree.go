package tree

//type TreeNode struct {
//    Val int
//    Left *TreeNode
//    Right *TreeNode
//}

type Element interface{}

type Queue interface {
    IsEmpty() bool
    Offer(ele Element) 
    Poll() Element
}

type ListQueue struct {
    elements []Element
}

func NewListQueue() *ListQueue{
   return &ListQueue{}
}

func (queue *ListQueue) IsEmpty() bool {
    return len(queue.elements) == 0
}

func (queue *ListQueue) Offer(e Element) {
    queue.elements = append(queue.elements, e) 
}

func (queue *ListQueue) Poll() Element {
    if queue.IsEmpty() {
        panic("queue is empty")
    }

    ele := queue.elements[0]
    queue.elements = queue.elements[1:]
    return ele
}

type LevelNode struct {
    node *TreeNode
    level int
}

func averageOfLevels(root *TreeNode) []float64 {
    if root == nil {
        return nil
    }
    var result []float64
    sum := 0
    nums := 0
    currentLevel := 1
    queue := NewListQueue()
    queue.Offer(&LevelNode{root, 1})

    for !queue.IsEmpty() {
        levelNode := queue.Poll().(*LevelNode)
        if levelNode.level == currentLevel {
            sum += levelNode.node.Val
            nums++
        }else {
            currentLevel = levelNode.level
            result = append(result, float64(sum) / float64(nums))
            sum = levelNode.node.Val
            nums = 1
        }

        if levelNode.node.Left != nil {
            queue.Offer(&LevelNode{levelNode.node.Left, currentLevel+1})
        }

        if  levelNode.node.Right != nil {
            queue.Offer(&LevelNode{levelNode.node.Right, currentLevel+1})
        }
    }

    result = append(result, float64(sum) / float64(nums))

    return result
}
