package main

import "fmt"

type Heap struct {
    count int
    nums []int //first ele empty 
}

func Construct() *Heap {
    heap := &Heap{}
    heap.count = 0
    heap.nums = []int{0}
    return heap
}

func (self *Heap)Insert(num int,maxHeap bool) {
    self.count++
    self.nums = append(self.nums, num)
    cursor := self.count
    //heapify from bottom 2 top
    for cursor > 1 {
	parent := cursor / 2
	if maxHeap {
	    if self.nums[parent] < self.nums[cursor] {
		swap(self.nums, parent, cursor)
		cursor = parent
	    }else {
		break
	    }
	}else {
	    if self.nums[parent] > self.nums[cursor] {
                swap(self.nums, parent, cursor)
                cursor = parent
            }else {
                break
            }
	}
    }
}

func (self *Heap) Replace(num int, maxHeap bool) int {
    toReturn := self.nums[1]
    self.nums[1] = num
    //heapify from top 2 bottom
    cursor := 1
    for cursor * 2 <= self.count {
	left, right := cursor * 2, cursor * 2 + 1
	if maxHeap {
	    maxIndex, max := cursor, self.nums[cursor]
	    if max < self.nums[left] {
		maxIndex = left
		max = self.nums[left]
	    }

	    if right <= self.count && max < self.nums[right] {
		maxIndex = right
		max = self.nums[right]
	    }

	    if maxIndex == cursor {
		return toReturn
	    }

	    swap(self.nums, cursor, maxIndex)
	    cursor = maxIndex
	}else {
	    minIndex, min := cursor, self.nums[cursor]
	    if min > self.nums[left] {
                minIndex = left
                min = self.nums[left]
            }

            if right <= self.count && min > self.nums[right] {
                minIndex = right
                min = self.nums[right]
            }

            if minIndex == cursor {
                return toReturn
            }

            swap(self.nums, cursor, minIndex)
            cursor = minIndex
	}
    }
    return toReturn
}

func swap(nums []int, a, b int) {
    nums[a], nums[b] = nums[b], nums[a]
}

type MedianFinder struct {
    count int
    maxHeap *Heap
    minHeap *Heap
}

func Constructor() MedianFinder {
    medianFinder := MedianFinder{}
    medianFinder.count = 0
    medianFinder.maxHeap = Construct()
    medianFinder.minHeap = Construct()
    return medianFinder
}

func (this *MedianFinder) AddNum(num int)  {
    this.count++
    fmt.Println("count:",this.count)
    if this.count % 2 == 1 {
	//add to maxHeap
	if len(this.minHeap.nums) > 1 &&  this.minHeap.nums[1] < num {
	    num = this.minHeap.Replace(num, false)
	}
	this.maxHeap.Insert(num, true)
    }else {
	//add to minHeap
	if len(this.maxHeap.nums) > 1 && this.maxHeap.nums[1] > num {
	    num = this.maxHeap.Replace(num, true)
	}
	this.minHeap.Insert(num, false)
    }

    fmt.Println("max: ", this.maxHeap.nums, " min: ", this.minHeap.nums)
}

func (this *MedianFinder) FindMedian() float64 {
    if this.count % 2 == 1 {
	return float64(this.maxHeap.nums[1]) * 1.0
    }

    return float64(this.maxHeap.nums[1] + this.minHeap.nums[1]) * 1.0 / 2
}

func main() {
    medianFinder := Constructor()
    medianFinder.AddNum(1)
    fmt.Println(medianFinder.FindMedian())
    medianFinder.AddNum(2)
    fmt.Println(medianFinder.FindMedian())
    medianFinder.AddNum(3)
    fmt.Println(medianFinder.FindMedian())
}
