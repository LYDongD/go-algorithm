package main

import "fmt"

func findOrder(numCourses int, prerequisites [][]int) []int {
	result := []int{}
	if numCourses == 0 {
		return result
	}

	if len(prerequisites) == 0 {
		for i := 0; i < numCourses; i++ {
			result = append(result, i)
		}
		return result
	}

	graph, inDegrees := generateGraph(numCourses, prerequisites)

	var queue []int
	//offer course with 0 degree
	for course, degree := range inDegrees {
		if degree == 0 {
			result = append(result, course)
			queue = append(queue, course)
		}
	}

	for len(queue) > 0 {
		preCourse := queue[0]
		postCourses := graph[preCourse]
		for _, course := range postCourses {
			inDegrees[course]--
			if inDegrees[course] == 0 {
				result = append(result, course)
				queue = append(queue, course)
			}
		}
		queue = queue[1:]
	}

	for _, degree := range inDegrees {
		if degree != 0 {
			return []int{}
		}
	}

	return result
}

func generateGraph(numCourses int, prerequisites [][]int) ([][]int, []int) {
	graph := make([][]int, numCourses)
	inDegrees := make([]int, numCourses)
	for row := 0; row < len(prerequisites); row++ {
		post := prerequisites[row][0]
		pre := prerequisites[row][1]
		graph[pre] = append(graph[pre], post)
		inDegrees[post]++
	}

	return graph, inDegrees
}

func main() {
	fmt.Println(findOrder(2, [][]int{{1, 0}}))
	fmt.Println(findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
}
