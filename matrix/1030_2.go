package main

import "fmt"

type Queue struct {
	nums [][]int
}

func (self *Queue) take() []int {
	if self.isEmpty() {
		panic("queue is empty!!")
	}

	head := self.nums[0]
	self.nums = self.nums[1:len(self.nums)]
	return head
}

func (self *Queue) offer(ele []int) {
	self.nums = append(self.nums, ele)
}

func (self *Queue) isEmpty() bool {
	return len(self.nums) == 0
}

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	if r0 < 0 || c0 < 0 || r0 > R || c0 > C {
		panic("paramete illegal")
	}

	queue := &Queue{nums: [][]int{{r0, c0}}}
	visited := make([][]bool, R)
	for index := range visited {
		visited[index] = make([]bool, C)
	}

	var result [][]int

	for !queue.isEmpty() {
		ele := queue.take()
		r, c := ele[0], ele[1]
		if r < 0 || c < 0 || r >= R || c >= C {
			continue
		}

		if visited[r][c] {
			continue
		}

		result = append(result, ele)
		visited[r][c] = true
		queue.offer([]int{r - 1, c})
		queue.offer([]int{r + 1, c})
		queue.offer([]int{r, c - 1})
		queue.offer([]int{r, c + 1})
	}

	return result
}

func main() {
	fmt.Println(allCellsDistOrder(1, 2, 0, 0))
	fmt.Println(allCellsDistOrder(2, 2, 0, 1))
	fmt.Println(allCellsDistOrder(2, 3, 1, 2))
}
