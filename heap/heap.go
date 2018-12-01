package main

import "fmt"

type Heap struct {
	a     []int
	n     int
	count int
}

//init heap
func NewHeap(capacity int) *Heap {
	heap := &Heap{}
	heap.n = capacity
	heap.a = make([]int, capacity+1)
	heap.count = 0
	return heap
}

//top-max heap -> heapify from down to up
func (heap *Heap) insert(data int) {
	//defensive
	if heap.count == heap.n {
		return
	}

	heap.count++
	heap.a[heap.count] = data

	//compare with parent node
	i := heap.count
	parent := i / 2
	for parent > 0 && heap.a[parent] < heap.a[i] {
		swap(heap.a, parent, i)
		i = parent
		parent = i / 2
	}
}

//heapfify from up to down
func (heap *Heap) removeMax() {

	//defensive
	if heap.count == 0 {
		return
	}

	//swap max and last
	swap(heap.a, 1, heap.count)
	heap.count--

	//comapre with left and right
	for i := 1; i <= heap.count/2; {

		maxIndex := i
		if heap.a[i] < heap.a[i*2] {
			maxIndex = i * 2
		}

		if i*2+1 <= heap.count && heap.a[maxIndex] < heap.a[i*2+1] {
			maxIndex = i*2 + 1
		}

		if maxIndex == i {
			break
		}

		swap(heap.a, i, maxIndex)
		i = maxIndex
	}

}

//swap two elements
func swap(a []int, i int, j int) {
	tmp := a[i]
	a[i] = a[j]
	a[j] = tmp
}

func main() {

	//	a := []int{1, 5, 7, 2, 4, 6}
	heap := NewHeap(10)
	heap.insert(1)
	heap.insert(5)
	heap.insert(7)
	heap.insert(4)
	heap.insert(2)
	heap.insert(6)

	fmt.Println(heap.a)
	fmt.Println("heap count: ", heap.count)
	heap.removeMax()
	fmt.Println(heap.a)
	heap.removeMax()
	fmt.Println(heap.a)
	heap.removeMax()
	fmt.Println(heap.a)
	heap.removeMax()
	fmt.Println(heap.a)
	heap.removeMax()
	fmt.Println(heap.a)
	heap.removeMax()
	fmt.Println(heap.a)
}
