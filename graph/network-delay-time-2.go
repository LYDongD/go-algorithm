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

func networkDelayTime(times [][]int, N int, K int) int {
	if len(times) == 0 {
		return -1
	}

	//构建图，邻接表法
	graph := configureGraph(times, N)

	//从K出发，记录到每个节点的最小距离，记录到hash
	distHash := make(map[int]int)
	//init hash
	for i := 1; i <= N; i++ {
		distHash[i] = int(^uint(0) >> 1)
	}

	//深度优先，计算从K出发可达定点的最小距离
	searchDistance(graph, distHash, K, 0)

	//求所有点最小距离中的最大距离
	result := 0
	for _, dist := range distHash {

		//如果有节点的最小距离没有更新，说明存在不可达顶点
		if dist == int(^uint(0)>>1) {
			return -1
		}
		if dist > result {
			result = dist
		}

	}

	return result
}

//计算从特定节点到其目标点的距离
func searchDistance(graph *Graph, distHash map[int]int, nodeVal int, currentDist int) {

	//累积距离已超过当前节点的最短距离，结束递归
	if currentDist >= distHash[nodeVal] {
		return
	}

	//先更新特定节点的最短距离
	distHash[nodeVal] = currentDist

	//从当前节点出发，累加到其目标节点的距离
	if graph.nodes[nodeVal] != nil {
		cursor := graph.nodes[nodeVal]
		for cursor != nil {
			searchDistance(graph, distHash, cursor.val, cursor.dist+currentDist)
			cursor = cursor.next
		}
	}
}

func main() {
	times := [][]int{{1, 2, 1}, {2, 1, 3}}
	fmt.Println(networkDelayTime(times, 2, 2))
}
