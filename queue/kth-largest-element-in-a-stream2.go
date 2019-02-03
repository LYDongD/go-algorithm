package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (this IntHeap) Len() int {
	return len(this)
}

func (this IntHeap) Less(i, j int) bool {
	return this[i] < this[j]
}

func (this IntHeap) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this *IntHeap) Push(x interface{}) {
	*this = append(*this, x.(int))
}

func (this *IntHeap) Pop() interface{} {

	old := *this
	n := len(old)
	last := old[n-1]
	*this = old[:n-1]
	return last
}

type KthLargest struct {
	Nums *IntHeap
	K    int
}

func Constructor(k int, nums []int) KthLargest {
	h := &IntHeap{}
	heap.Init(h)
	for i := 0; i < len(nums); i++ {
		heap.Push(h, nums[i])
	}

	for len(*h) > k {
		heap.Pop(h)
	}

	return KthLargest{h, k}
}

func (this *KthLargest) Add(val int) int {
	if len(*this.Nums) < this.K {
		heap.Push(this.Nums, val)
	} else {
		heap.Push(this.Nums, val)
		heap.Pop(this.Nums)
	}

	return (*this.Nums)[0]
}
