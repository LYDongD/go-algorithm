package main

import "fmt"

func findMinHeightTrees(n int, edges [][]int) []int {

    //defensive coding
    if n < 0 {
	return []int{}
    }

    if n == 1 || len(edges) == 0 {
	return []int{0}
    }

    //build graph by adjacent table
    graph, degrees := initGraphAndDegree(n, edges)
    fmt.Println("graph:", graph, "degree", degrees)

    //BFS shrik from leaf
    queue := []int{}
    for vertix, degree := range degrees {
	if degree == 1 {
	    queue = append(queue, vertix)
	}
    }

    result := []int{}
    for len(queue) > 0 {

	//cycle leaf shrink
	leafCount := len(queue)
	result = []int{}
	for i := 0; i < leafCount; i++ {
	    head := queue[0]
	    queue = queue[1:]
	    result = append(result, head)
	    //update degree from graph and offer vertix with 1 degree(next leaf shrink)
	    edgeVertix := graph[head]
	    for i := 0; i < len(edgeVertix); i++ {
		degrees[edgeVertix[i]]--
		if degrees[edgeVertix[i]] == 1 {
		    queue = append(queue, edgeVertix[i])
		}
	    }
	}
    }

    return result
}

func initGraphAndDegree(n int, edges [][]int) ([][]int, []int)  {
    graph := make([][]int,n)
    degree := make([]int, n)
    for i := 0; i < len(edges); i++ {
	vertixA, vertixB := edges[i][0], edges[i][1]
	graph[vertixA] = append(graph[vertixA], vertixB)
	graph[vertixB] = append(graph[vertixB], vertixA)
	degree[vertixA]++
	degree[vertixB]++
    }

    return graph, degree
}


func main() {
    fmt.Println(findMinHeightTrees(6, [][]int{{0,1},{0,2},{0,3}, {3,4}, {4,5}}))
    fmt.Println(findMinHeightTrees(4, [][]int{{1,0},{1,2},{1,3}}))
    fmt.Println(findMinHeightTrees(6, [][]int{{0,3},{1,3},{2,3}, {4,3}, {5,4}}))
    fmt.Println(findMinHeightTrees(0, [][]int{}))
}
