package main

import "fmt"

type KthLargest struct {
	head   *Node
	tail   *Node
	isFull bool
}

type Node struct {
	val  int
	next *Node
	pre  *Node
}

func Constructor(k int, nums []int) KthLargest {

	//1 sort
	quickSort(nums, 0, len(nums)-1)

	//2 construct LinkedList
	head, tail, isFull := initLinkedList(nums, k)

	return KthLargest{head, tail, isFull}
}

func quickSort(nums []int, start int, end int) {
	if start >= end {
		return
	}

	q := partion(nums, start, end)
	quickSort(nums, start, q-1)
	quickSort(nums, q+1, end)
}

func partion(nums []int, start int, end int) int {
	//select a pivot
	pivot := nums[end]

	//find the index to split
	q := start
	for i := start; i < end; i++ {
		if nums[i] > pivot {
			swap(nums, q, i)
			q++
		}
	}

	swap(nums, q, end)
	return q
}

func swap(nums []int, a int, b int) {
	nums[a], nums[b] = nums[b], nums[a]
}

func initLinkedList(nums []int, k int) (*Node, *Node, bool) {

	if len(nums) == 0 {
		return nil, nil, false
	}

	var head *Node
	var tail *Node
	if len(nums) == 1 {
		head = &Node{nums[0], nil, nil}
		tail = head
		return head, tail, len(nums) >= k
	}

	for i := 0; i < k-1; i++ {
		if head == nil {
			head = &Node{nums[i], nil, nil}
			tail = head
		} else {
			node := &Node{nums[i], nil, tail}
			tail.next = node
			tail = tail.next
		}
	}

	if len(nums) >= k {
		tail.next = &Node{nums[k-1], nil, tail}
		tail = tail.next
	}

	return head, tail, len(nums) >= k
}

func (this *KthLargest) Add(val int) int {

	if this.head == nil {
		this.head = &Node{val, nil, nil}
		this.tail = this.head
		this.isFull = true
		return val
	}

	//insert
	cursor := this.head
	for cursor != nil {
		//insert
		if val > cursor.val {
			node := &Node{val, cursor, cursor.pre}
			if cursor.pre == nil {
				this.head = node
			} else {
				cursor.pre.next = node
			}

			cursor.pre = node
			break
		}
		cursor = cursor.next
	}

	if cursor == nil {
		node := &Node{val, nil, this.tail}
		this.tail.next = node
		this.tail = this.tail.next
	}

	if !this.isFull {
		this.isFull = true
		return this.tail.val
	}

	this.tail = this.tail.pre
	this.tail.next = nil

	return this.tail.val
}

func printHead(head *Node) {
	tmp := head
	for tmp != nil {
		fmt.Print(tmp.val)
		tmp = tmp.next
	}
	fmt.Println("")
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
func main() {
	nums := []int{-2}
	//quickSort(nums, 0, len(nums)-1)
	//fmt.Println(nums)

	//head, tail, length := initLinkedList(nums, 4)
	//fmt.Println(length)
	//for head != nil {
	//	fmt.Print(head.val)
	//	head = head.next
	//}

	//fmt.Println("")

	//for tail != nil {
	//	fmt.Print(tail.val)
	//	tail = tail.pre
	//}
	//fmt.Println("")

	kthLargest := Constructor(1, nums)
	fmt.Println(kthLargest.Add(-3))
	fmt.Println(kthLargest.Add(0))
	fmt.Println(kthLargest.Add(2))
	fmt.Println(kthLargest.Add(-1))
	fmt.Println(kthLargest.Add(4))
}
