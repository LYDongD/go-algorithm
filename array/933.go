package main

import "fmt"

type Queue struct {
    nums []int
}

func (this *Queue) size() int {
    return len(this.nums)
}

func (this *Queue) add(num int) {
    this.nums = append(this.nums, num)
}

func (this *Queue) peek() int {
    if this.size() > 0 {
	return this.nums[0]
    }

    panic("queue is mempty")
}

func (this *Queue) poll() {
    if this.size() > 0 {
	this.nums = this.nums[1:]
	return
    }

    panic("queue is mempty")
}

type RecentCounter struct {
    queue *Queue
}


func Constructor() RecentCounter {
    queue := &Queue{nums : []int{}}
    return RecentCounter{queue : queue}
}


func (this *RecentCounter) Ping(t int) int {
    this.queue.add(t)
    for this.queue.peek() < t - 3000 {
	this.queue.poll()
    }

    return this.queue.size()
}



func main() {
    counter := Constructor()
    fmt.Println(counter.Ping(1))
    fmt.Println(counter.Ping(100))
    fmt.Println(counter.Ping(3001))
    fmt.Println(counter.Ping(3002))
}
