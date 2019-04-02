package main

import "fmt"

type RecentCounter struct {
    pingList []int
}


func Constructor() RecentCounter {
    return RecentCounter{pingList : []int{}}
}


func (this *RecentCounter) Ping(t int) int {
    if len(this.pingList) == 0 {
	this.pingList = append(this.pingList, t)
	return 1
    }

    lowerBound := t - 3000
    upperBound := t

    newPingList := []int{}
    count := 0
    for _, ping := range this.pingList {
	if ping  >= lowerBound && ping <= upperBound {
	    newPingList = append(newPingList, ping)
	    count++
	}
    }
    fmt.Println(newPingList)
    this.pingList = newPingList
    return count
}

func main() {
    counter := Constructor()
    fmt.Println(counter.Ping(1))
    fmt.Println(counter.pingList)
    fmt.Println(counter.Ping(100))
    fmt.Println(counter.pingList)
    fmt.Println(counter.Ping(3001))
    fmt.Println(counter.pingList)
    fmt.Println(counter.Ping(3002))
    fmt.Println(counter.pingList)
}
