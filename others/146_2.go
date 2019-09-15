package main

import "fmt"

type LRUCache struct {
    dic map[int]*ListNode
    head *ListNode
    tail *ListNode
    capacity int
}

type ListNode struct {
    key int
    val int
    next *ListNode
    prev *ListNode
}

func Constructor(capacity int) LRUCache {
    head, tail := &ListNode{},&ListNode{}
    head.next = tail
    tail.prev = head
    return LRUCache{
	dic : make(map[int]*ListNode),
	head : head,
	tail : tail,
	capacity : capacity,
    }
}

func (this *LRUCache) Get(key int) int {
    node, ok := this.dic[key]
    if !ok {
	return -1
    }

    this.moveToHead(node)
    return node.val
}

func (this *LRUCache) Put(key int, value int)  {
    node, ok := this.dic[key]
    if ok {
	node.val = value
	this.moveToHead(node)
	return
    }

    if len(this.dic) == this.capacity {
	lastNode := this.tail.prev
	this.removeNode(lastNode)
	delete(this.dic, lastNode.key)
    }

    newNode := &ListNode{key:key, val : value}
    this.addToHead(newNode)
    this.dic[key] = newNode
}


func (this *LRUCache) moveToHead(node *ListNode) {
    this.removeNode(node)
    this.addToHead(node)
}

func (this *LRUCache) addToHead(node *ListNode) {
    node.next = this.head.next
    this.head.next.prev = node
    this.head.next=node
    node.prev = this.head
}

func (this *LRUCache) removeNode(node *ListNode) {
    node.prev.next = node.next
    node.next.prev = node.prev
    node.next = nil
    node.prev = nil
}


func main() {
    fmt.Println("hah")
}
