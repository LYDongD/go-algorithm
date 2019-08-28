package main

import "fmt"

type Heap struct {
	nums     []int
	capacity int
	count    int
}

func Constructor(k int) *Heap {
	heap := &Heap{}
	heap.capacity = k
	heap.count = 0
	heap.nums = make([]int, k+1)
	return heap
}

func (self *Heap) heapipyDownToUp() {
	child := self.count
	parent := child / 2
	for parent >= 1 {
		if self.nums[parent] <= self.nums[child] {
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
	rightChild := parent*2 + 1
	minIndex := parent
	for {
		if leftChild <= self.count && self.nums[parent] > self.nums[leftChild] {
			minIndex = leftChild
		}

		if rightChild <= self.count && self.nums[minIndex] > self.nums[rightChild] {
			minIndex = rightChild
		}

		if minIndex == parent {
			break
		}

		swap(self.nums, parent, minIndex)
		parent = minIndex
		leftChild = parent * 2
		rightChild = parent*2 + 1
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

func (self *Heap) top() int {
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
			if heap.top() > num {
				continue
			}
			heap.removeTop()
		}
		heap.insert(num)
		fmt.Println(heap.nums)
	}
	return heap.top()
}

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	//fmt.Println([]int{3, 2, 3, 1, 2, 4, 5, 5, 6})
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}
