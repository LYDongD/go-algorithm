package main

import "fmt"

type NumArray struct {
    segmentTree []int
}


func Constructor(nums []int) NumArray {
    numArr := NumArray{}
    numArr.segmentTree = buildTree(nums)
    return numArr
}

func buildTree(nums []int) []int {
    n := len(nums)
    tree := make([]int, 2 * n)

    //init leaf node
    for i := 0; i < n; i++ {
	tree[i + n] = nums[i]
    }

    //generate parent
    for i := n - 1; i > 0; i-- {
	tree[i] = tree[i * 2] + tree[i * 2 + 1]
    }

    return tree
}


func (this *NumArray) Update(i int, val int)  {
    n := len(this.segmentTree) / 2
    this.segmentTree[i+n] = val
    i += n
    for i > 0 {
	parent := i / 2
	//left child
	if i % 2 == 0 {
	    this.segmentTree[parent] = this.segmentTree[i] + this.segmentTree[i+1]
	}else { //right child
	    this.segmentTree[parent] = this.segmentTree[i-1] + this.segmentTree[i]
	}
	i = parent
    }
}


func (this *NumArray) SumRange(i int, j int) int {
    if i > j {
	panic("illegal index bound")
    }
    n := len(this.segmentTree) / 2
    sum := 0
    i += n
    j += n
    for i <= j {
	if  i  % 2 == 1 {
	    sum += this.segmentTree[i]
	    i++
	}

	if  j  % 2 == 0 {
	    sum += this.segmentTree[j]
	    j--
	}

	i = i / 2
	j = j / 2
    }

    return sum
}

func main() {
    obj := Constructor([]int{1,3,5})
    fmt.Println(obj.SumRange(0,2))
    obj.Update(1,2)
    fmt.Println(obj.SumRange(0,2))

}


