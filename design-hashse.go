package main

type MyHashSet struct {
	buckets []int
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{make([]bool, 1000000)}
}

func (this *MyHashSet) Add(key int) {
	this.buckets[key] = true
}

func (this *MyHashSet) Remove(key int) {
	this.buckets[key] = false
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	return this.buckets[key]
}
