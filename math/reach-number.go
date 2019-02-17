package main

import "fmt"

func reachNumber(target int) int {
    if target == 0 {
	return 0
    }

    if target < 0 {
	target = -target
    }

    //往前走，直到超过target
    k := 0
    sum := 0
    for sum < target {
	k++ 
	sum += k
    }

    //如果超过距离为偶数，调整回退原来中的某一步即可，不影响总步数
    delta := sum - target
    if delta % 2 == 0 {
	return k
    }

    //如果超过距离为奇数，根据当前k的奇偶性向前走1步或2步使超过距离到达偶数
    if k % 2 == 0 {
	return k + 1
    }

    return k + 2

}

func main() {
    fmt.Println(reachNumber(3))
    fmt.Println(reachNumber(2))
}
