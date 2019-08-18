package main

import "fmt"

func canFinish(numCourses int, prerequisites [][]int) bool {
	adjArr := generateAdjArr(numCourses, prerequisites)
	checkedDic := make(map[int]int)
	for i := 0; i < len(adjArr); i++ {
		hasCircle := dfsCheck(adjArr, checkedDic, i)
		if hasCircle {
			return false
		}
	}

	return true
}

func generateAdjArr(numCourses int, prerequisites [][]int) [][]int {
	adjArr := make([][]int, numCourses)
	for row := 0; row < len(prerequisites); row++ {
		vertex := prerequisites[row][1]
		adjArr[vertex] = append(adjArr[vertex], prerequisites[row][0])
	}

	return adjArr
}

func dfsCheck(adjArr [][]int, checkedDic map[int]int, i int) bool {
	if checkedDic[i] > 0 {
		return true
	}

	checkedDic[i]++
	toVertexArr := adjArr[i]
	for k := 0; k < len(toVertexArr); k++ {
		hasCircle := dfsCheck(adjArr, checkedDic, toVertexArr[k])
		if hasCircle {
			return true
		}
	}

	checkedDic[i]--
	return false
}
func main() {
	fmt.Println(canFinish(2, [][]int{{1, 0}}))
	fmt.Println(canFinish(2, [][]int{{1, 0}, {0, 1}}))
	fmt.Println(canFinish(3, [][]int{{0, 1}, {0, 2}, {1, 2}}))
}
