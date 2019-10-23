package main

import "fmt"

type NumArray struct {
    nums []int
    sums []int
}


func Constructor(nums []int) NumArray {
    sums := make([]int, len(nums))
    sum := 0
    for i := 0; i < len(nums); i++ {
	sum += nums[i]
	sums[i] = sum
    }

    numArr := NumArray{}
    numArr.sums = sums
    numArr.nums = nums
    return numArr
}


func (this *NumArray) Update(i int, val int)  {
    origin := this.nums[i]
    delta := val - origin
    this.nums[i] = val
    for k := i; k < len(this.sums); k++ {
	this.sums[k] = this.sums[k] + delta
    }
}


func (this *NumArray) SumRange(i int, j int) int {
    return this.sums[j] - this.sums[i] + this.nums[i]
}

func main() {
    obj := Constructor([]int{1,3,5})
    fmt.Println(obj.SumRange(0,2))
    obj.Update(1,2)
    fmt.Println("nums: ", obj.nums, "sum: ", obj.sums)
    fmt.Println(obj.SumRange(0,2))
}
