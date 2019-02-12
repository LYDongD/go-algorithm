package main

import "fmt"

func networkDelayTime(times [][]int, N int, K int) int {

	//configure grah
	graph := configureGraph(times, N)

	//prepare vertix array
	vertixs := make([]*Vertix, N+1)
	for i := 1; i < N+1; i++ {
		vertixs[i] = &Vertix{i, int(^uint(0) >> 1)}
	}

	//Dijkstra find shortest path
	delayTime := 0
	for i = 1; i < N; i++ {
		if i != K {
			time := dijkstra(graph, vertixs, K, i)

			//unreeach
			if time == -1 {
				return -1
			}

			//reach and get the max time for all
			if time > delayTime {
				delayTime = time
			}
		}
	}
	return delayTime
}

type Graph struct {
	V     int     //顶点个数
	Nodes []*Node //链表头节点数组
}

type Node struct {
	Val  int
	Time int
	Next *Node
}

type Vertix struct {
	id        int
	leastTime int
}

func configureGraph(times [][]int, v int) *Graph {
	graph := &Graph{v, make([]*Node, v+1)}
	for i := 0; i < len(times); i++ {
		source := times[i][0]
		target := times[i][1]
		time := times[i][2]
		if graph.Nodes[source] == nil {
			graph.Nodes[source] = &Node{target, time, nil}
		} else {
			cursor := graph.Nodes[source]
			for cursor.Next != nil {
				cursor = cursor.Next
			}
			cursor.Next = &Node{target, time, nil}
		}
	}

	return graph
}

//Dijkstra
func dijkstra(graph []*Node, vertixs []*Vertix.s, t int) int {

	//广度优先，将边相关的节点加入优先级队列
	priorityQueue := &PriorityQueue{graph.V}
	priorityQueue.Add(vertixs[s])

	//记录已经入队的队列
	inQueue := make([]bool, graph.V+1)

	for !priorityQueue.IsEmpty() {
		minVertix := priorityQueue.Poll()
		if minVertix.id == t {
			return minVertix.leastTime
		}

		//遍历以当前节点为起始节点的边
		node := graph.Nodes[K]
		nextVertix := vertixs[node.Val]
		for node != nil {

			//计算路径长度,如果当前路径较短，则将该边目标节点作可选节点
			if minVertix.leastTime+node.Time < nextVertix.leastTime {
				//更新目标节点的最短路径
				nextVertix.leastTime = minVertix.leastTime + node.Time

				//加入队列或更新队列
				if inQueue[nextVertix.id] {
					priorityQueue.Update(nextVertix)
				} else {
					priorityQueue.Add(nextVertix)
					inQueue[nextVertix.id]
				}
			}
			node = node.Next
		}
	}

}

func main() {

	times := [][]int{{1, 2, 1}, {2, 3, 2}, {1, 3, 2}}
	fmt.Println(networkDelayTime(times, 3, 1))
}
