package main

import "fmt"

type LRUCache struct {
    dic map[int]int
    orderedKeys []int
    capacity int
}

func Constructor(capacity int) LRUCache {
    return LRUCache{capacity:capacity,
    orderedKeys:[]int{},
    dic : make(map[int]int)}
}

func (this *LRUCache) Get(key int) int {
    val, ok := this.dic[key]
    if !ok {
	return -1
    }

    i := findKeyPosition(this.orderedKeys, key)
    //move key to the tail
    this.orderedKeys = append(this.orderedKeys[:i], this.orderedKeys[i+1:]...)
    this.orderedKeys = append(this.orderedKeys, key)
    return val
}

func (this *LRUCache) Put(key int, value int)  {

    val := this.Get(key)
    //get operation has update order
    if val != -1 {
	this.dic[key] = value
	return
    }

    if len(this.orderedKeys) == this.capacity {
	//remove dic
	delete(this.dic, this.orderedKeys[0])
	//remove the head
	this.orderedKeys = append(this.orderedKeys[:0], this.orderedKeys[1:]...)
    }

    //add new 
    this.dic[key] = value
    //add to orderedKeys
    this.orderedKeys = append(this.orderedKeys, key)
}

func findKeyPosition(orderedKeys []int, key int) int {
    for i, orderKey := range orderedKeys {
	if orderKey == key {
	    return i
	}
    }

    return -1
}

func main() {
    fmt.Println("haha")
}
