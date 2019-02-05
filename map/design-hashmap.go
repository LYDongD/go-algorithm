package main

import "fmt"

type MyHashMap struct {
	nums []*Node
	size int
}

//locate key position
func hash(key int, capacity int) int {
	return key % capacity
}

//linked list
type Node struct {
	key   int
	value int
	next  *Node
}

func (self *MyHashMap) addNode(head *Node, key int, value int) *Node {

	cursor := head
	for cursor.next != nil {
		if cursor.key == key {
			cursor.value = value
		}

		cursor = cursor.next
	}

	if cursor.key == key {
		cursor.value = value
	} else {
		cursor.next = &Node{key, value, nil}
		self.size++
	}

	return cursor
}

func (self *MyHashMap) removeNode(head *Node, key int) *Node {
	if head == nil {
		return nil
	}

	cursor := head
	last := cursor

	for cursor != nil {
		if cursor.key == key {
			if cursor == head {
				head = cursor.next
			} else {
				last.next = cursor.next
			}
			cursor = nil
			self.size--
			return head
		}

		last = cursor
		cursor = cursor.next

	}

	return head
}

func (self *MyHashMap) getNodeValue(head *Node, key int) int {
	if head == nil {
		return -1
	}

	cursor := head
	for cursor != nil {
		if cursor.key == key {
			return cursor.value
		}

		cursor = cursor.next
	}

	return -1
}

/** Initialize your data structure here. */
func Constructor() MyHashMap {
	return MyHashMap{make([]*Node, 4), 0}
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	index := hash(key, len(this.nums))
	if this.nums[index] == nil {
		this.nums[index] = &Node{key, value, nil}
		this.size++
	} else {
		this.addNode(this.nums[index], key, value)
	}

	fmt.Println("current size: ", this.size)

	if this.size == len(this.nums) {
		this.size = 0
		this.enlarge()
	}

}

func (self *MyHashMap) enlarge() {
	newNums := make([]*Node, len(self.nums)*2)
	nums := self.nums
	self.nums = newNums
	for _, node := range nums {
		for node != nil {
			self.Put(node.key, node.value)
			node = node.next
		}
	}
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	index := hash(key, len(this.nums))
	return this.getNodeValue(this.nums[index], key)

}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	index := hash(key, len(this.nums))
	this.nums[index] = this.removeNode(this.nums[index], key)
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */

func main() {
	obj := Constructor()
	obj.Put(1, 1)
	obj.Put(2, 2)
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(3))
	obj.Put(2, 1)
	fmt.Println(obj.Get(2))
	obj.Remove(2)
	fmt.Println(obj.Get(2))
}
