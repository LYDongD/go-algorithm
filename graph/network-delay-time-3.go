package main

import "fmt"

type Graph struct {
	v     int     //顶点个数
	nodes []*Node //链表头节点数组
}

type Node struct {
	val  int
	dist int
	next *Node
}

//构建grap，单链表数组
func configureGraph(times [][]int, vertixCount int) *Graph {
	graph := &Graph{vertixCount, make([]*Node, vertixCount+1)}
	for i := 0; i < len(times); i++ {
		source := times[i][0]
		target := times[i][1]
		dist := times[i][2]
		if graph.nodes[source] == nil {
			graph.nodes[source] = &Node{target, dist, nil}
		} else {
			cursor := graph.nodes[source]
			for cursor.next != nil {
				cursor = cursor.next
			}
			cursor.next = &Node{target, dist, nil}
		}
	}

	return graph
}

//顶点
type Vertix struct {
	id   int //编号
	dist int //最短距离
}

func networkDelayTime(times [][]int, N int, K int) int {
	if len(times) == 0 {
		return -1
	}

	//构建图，邻接表法
	graph := configureGraph(times, N)

	//从K出发， 记录其他顶点到起点的最短距离，顶点自己记录
	vertixs := make([]*Vertix, graph.v+1)

	//初始化最短距离为最大值, 哨兵处理起始元素
	vertixs[0] = &Vertix{0, 0}
	for i := 1; i <= N; i++ {
		vertixs[i] = &Vertix{i, int(^uint(0) >> 1)}
	}

	//Dijkstra，使用优先级队列遍历节点实现搜索
	dijkstra(graph, vertixs, K)

	//求所有点最小距离中的最大距离
	result := 0
	for _, vertix := range vertixs {

		//如果有节点的最小距离没有更新，说明存在不可达顶点
		if vertix.dist == int(^uint(0)>>1) {
			return -1
		}

		if vertix.dist > result {
			result = vertix.dist
		}
	}
	return result
}

func dijkstra(graph *Graph, vertixs []*Vertix, start int) {

	//通过优先级队列遍历所有连接点，并更新其最小距离
	queue := &PriorityQueue{}

	//加入起始节点，起始节点的最短距离为
	startVertix := &Vertix{start, 0}
	vertixs[start] = startVertix

	queue.add(startVertix)

	//记录已入队的节点，避免重复入队
	inQueue := make([]bool, graph.v+1)

	for !queue.isEmpty() {
		vertix := queue.poll()
		//遍历其所有邻接节点，更新最短距离
		edgeNode := graph.nodes[vertix.id]
		for edgeNode != nil {

			nextVertix := vertixs[edgeNode.val]

			//如果比领接节点原本的最小距离小，则更新，否则说明已经访问过获取相对较小的距离，不需要继续处理该节点
			if nextVertix.dist > vertix.dist+edgeNode.dist {
				nextVertix.dist = vertix.dist + edgeNode.dist

				//如果节点已经在队列中，则只需要更新dist并重新排序
				if inQueue[nextVertix.id] {
					queue.update(nextVertix)
				} else { //加入优先级队列
					queue.add(nextVertix)
					inQueue[nextVertix.id] = true
				}
			}
			edgeNode = edgeNode.next
		}
	}
}

type VertixNode struct {
	id   int
	dist int
	next *VertixNode
}

type PriorityQueue struct {
	head *VertixNode
	size int
}

func (self *PriorityQueue) add(vertix *Vertix) {

	//从头部添加
	if self.head == nil || vertix.dist < self.head.dist {
		newHead := &VertixNode{vertix.id, vertix.dist, self.head}
		self.head = newHead
		self.size++
	} else {
		cursor := self.head
		last := cursor
		for cursor != nil {

			if cursor.dist > vertix.dist {
				newNode := &VertixNode{vertix.id, vertix.dist, cursor}
				last.next = newNode
				self.size++
				return
			}

			last = cursor
			cursor = cursor.next
		}

		//从尾部添加
		last.next = &VertixNode{vertix.id, vertix.dist, cursor}
		self.size++
	}

}

func (self *PriorityQueue) update(vertix *Vertix) {

	//更新头节点
	if vertix.id == self.head.id {
		self.poll()
		self.add(vertix)
		return
	}

	cursor := self.head
	last := cursor
	for cursor != nil {

		if cursor.id == vertix.id {
			//删除并重新插入
			last.next = cursor.next
			self.size--
			self.add(vertix)
		}

		last = cursor
		cursor = cursor.next
	}
}

func (self *PriorityQueue) poll() *Vertix {

	//从头部弹出
	if self.head != nil {
		vertix := &Vertix{self.head.id, self.head.dist}
		self.head = self.head.next
		self.size--
		return vertix
	}

	return nil
}

func (self *PriorityQueue) isEmpty() bool {
	return self.size == 0
}

func main() {
	times := [][]int{{1, 2, 1}, {2, 1, 3}}
	fmt.Println(networkDelayTime(times, 2, 2))
}
