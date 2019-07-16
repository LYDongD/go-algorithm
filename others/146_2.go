package main

import "fmt"

type LRUCache struct {
    dic map[int]*ListNode
    head *ListNode
    tail *ListNode
    capacity int
}

func ListNode struct {
    val int
    next *ListNode
    prev *ListNode
}

func Constructor(capacity int) LRUCache {
    return &LRUCache{
	dic : make(map[int]*ListNode)
	head : &ListNode{}
	tail : &ListNode{}
	capacity : capacity
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

    newNode := &ListNode{val : value}
    addNode(newNode)
    dic[key] = newNode
}


func (this *LRUCache) moveToHead(node *ListNode) {
    this.removeNode(node)
    node.next = head.next
    head.next.prev = node
    head.next=node
    node.prev = head
}

func (this *LRUCache) removeNode(node *ListNode) {
    node.prev.next = node.next
    node.next.prev = node.prev
    node.next = nil
    node.prev = nil
}
