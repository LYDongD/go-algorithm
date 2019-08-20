package main

import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {
	//key -> need(source point), value -> want(dest points), ajacent list
	graph := make(map[int][]int)
	//index -> want, value -> source points
	dependencies := make([]int, numCourses)

	//init
	for row := 0; row < len(prerequisites); row++ {
		need := prerequisites[row][1]
		want := prerequisites[row][0]
		if _, ok := graph[need]; ok {
			graph[need] = append(graph[need], want)
		} else {
			graph[need] = []int{want}
		}

		dependencies[want]++
	}

	//use queue to search in BFS, start from pure source point
	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if dependencies[i] == 0 {
			queue = append(queue, i)
		}
	}

	//there should be at least one pure source point
	if len(queue) == 0 {
		return false
	}

	for len(queue) > 0 {
		head := queue[0]
		wants := graph[head]
		for k := 0; k < len(wants); k++ {
			dependencies[wants[k]]--
			if dependencies[wants[k]] == 0 {
				queue = append(queue, wants[k])
			}
		}

		queue = queue[1:]
	}

	for _, count := range dependencies {
		if count != 0 {
			return false
		}
	}

	return true

}

func main() {
	fmt.Println("vim go")
}
