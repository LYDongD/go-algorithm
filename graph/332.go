package main

import "fmt"
import "container/heap"

type Interface interface {
    Less(ohter interface{}) bool
}

type sorter []Interface

func (s *sorter) Push(x interface{}){
    *s = append(*s, x.(Interface))
}

func (s *sorter) Pop() interface{} {
    n := len(*s)
    if n > 0 {
	x := (*s)[n-1]
	*s = (*s)[0:n-1]
	return x
    }

    return nil
}

func (s *sorter) Len() int {
	return len(*s)
}

func (s *sorter) Less(i, j int) bool {
	return (*s)[i].Less((*s)[j])
}

func (s *sorter) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

// Define priority queue struct
type PriorityQueue struct {
	s *sorter
}

func New() *PriorityQueue {
	q := &PriorityQueue{s: new(sorter)}
	heap.Init(q.s)
	return q
}

func (q *PriorityQueue) Push(x Interface) {
	heap.Push(q.s, x)
}

func (q *PriorityQueue) Pop() Interface {
	return heap.Pop(q.s).(Interface)
}

func (q *PriorityQueue) Top() Interface {
	if len(*q.s) > 0 {
		return (*q.s)[0].(Interface)
	}
	return nil
}

func (q *PriorityQueue) Fix(x Interface, i int) {
	(*q.s)[i] = x
	heap.Fix(q.s, i)
}

func (q *PriorityQueue) Remove(i int) Interface {
	return heap.Remove(q.s, i).(Interface)
}

func (q *PriorityQueue) Len() int {
	return q.s.Len()
}

type Node struct {
    value string
}

func (this *Node)Less(other interface{}) bool {
    return this.value < other.(*Node).value
}

var result []string

func findItinerary(tickets [][]string) []string {

    result = []string{}
    if len(tickets) == 0 {
	return result
    }

    //init graph with priority queue
    graph := make(map[string]*PriorityQueue)
    for _, ticket := range tickets {
	from, to := ticket[0], ticket[1]
	if graph[from] == nil {
	    graph[from] = New()
	}
	queue := graph[from]
	queue.Push(&Node{to})
	graph[from] = queue
    }

    //dfs search graph 
   dfs("JFK", graph)
   return result
}

func dfs(from string,  graph map[string]*PriorityQueue) {
    toPriority := graph[from]
    for toPriority != nil && toPriority.Len() > 0 {
	dfs(toPriority.Pop().(*Node).value,  graph)
    }

    result = append([]string{from}, result...)
}

func main() {
    fmt.Println(findItinerary([][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}}))
    fmt.Println(findItinerary([][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}}))
    fmt.Println(findItinerary([][]string{{"JFK","KUL"}, {"JFK","NRT"}, {"NRT","JFK"}}))
}
