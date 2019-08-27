package main

import "fmt"


type Heap struct {
    nums []int
    capacity int
    count int
}

func Constructor(k int) *Heap{
    heap := &Heap{}
    heap.capacity = k
    heap.count = 0
    heap.nums = make([]int, k + 1)
    return heap
}

func (self *Heap) heapipyDownToUp() {
    child := self.count
    parent := child / 2
    for parent >= 1 {
	if self.nums[parent] > self.nums[child] {
	    break
	}
	swap(self.nums, parent, child)
	child = parent
	parent = child / 2
    }
}

func (self *Heap) heapifyUpToDown() {
    parent := 1
    leftChild := parent * 2
    rightChild := parent * 2 + 1
    maxIndex := parent
    for leftChild <= self.count {
	if self.nums[parent] < self.nums[leftChild] {
	    maxIndex = leftChild
	}

	if self.nums[maxIndex] < self.nums[rightChild] {
	    maxIndex = rightChild
	}

	swap(self.nums, parent, maxIndex)
	parent = maxIndex
	leftChild = parent * 2
	rightChild = parent * 2 + 1
    }
}

func (self *Heap) insert(num int) {
    self.count++
    self.nums[self.count] = num
    self.heapipyDownToUp()
}

func (self *Heap) removeTop() {
    swap(self.nums, 1, self.count)
    self.count--
    self.heapifyUpToDown()
}

func (self *Heap) top() int{
    return self.nums[1]
}

func (self *Heap) isFull() bool {
    return self.capacity == self.count
}

func swap(nums []int, a, b int) {
    nums[a], nums[b] = nums[b], nums[a]
}

func findKthLargest(nums []int, k int) int {
    heap := Constructor(k)
    for _, num := range nums {
	if heap.isFull() {
	    heap.removeTop()
	}
	heap.insert(num)
    }
    return heap.top()
}

func main() {
    fmt.Println(findKthLargest([]int{3,2,1,5,6,4}, 2))
}
