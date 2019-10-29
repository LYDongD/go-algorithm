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

func swap(nums []int, a, b int) {
	nums[a], nums[b] = nums[b], nums[a]
}

func (self *Heap) insert(num int) {
	self.count++
	self.nums[self.count] = num
	self.heapifyDownToUp() //route the least one to the top
}

func (self *Heap) removeTop() {
	swap(self.nums, 1, self.count)
	self.count--
	self.heapifyUpToDown() //router the leaset one to the top
}

func (self *Heap) top() int {
	return self.nums[1]
}

func (self *Heap) isFull() bool {
	return self.capacity == self.count
}

func findKthLargest(nums []int, k int) int {
	heap := Constructor(k)
	for _, num := range nums {
		if heap.isFull() {
			if num < heap.top() {
				continue
			}

			heap.removeTop()
		}

		heap.insert(num)
	}

	return heap.top()
}

func (self *Heap) heapifyDownToUp() {
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

func wiggleSort(nums []int) {
	median := findKthLargest(nums, (len(nums)+1)/2)
	n := len(nums)
	i, left, right := 0, 0, n-1
	for i <= right {
		if nums[checkIndex(i, n)] > median {
			swap(nums, checkIndex(left, n), checkIndex(i, n))
			left++
			i++
		} else if nums[checkIndex(i, n)] < median {
			swap(nums, checkIndex(right, n), checkIndex(i, n))
			right--
		} else {
			i++
		}
	}

}

func checkIndex(i, n int) int {
	return (i*2 + 1) % (n | 1)
}

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
	nums := []int{1, 5, 1, 1, 6, 4}
	wiggleSort(nums)
	fmt.Println(nums)

	nums2 := []int{1, 3, 2, 2, 3, 1}
	wiggleSort(nums2)
	fmt.Println(nums2)
}
