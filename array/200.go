package main

import "fmt"

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	rowCount, colCount := len(grid), len(grid[0])

	islandCount := 0

	visited := makeVisitedGrid(rowCount, colCount)
	for row := 0; row < rowCount; row++ {
		for col := 0; col < colCount; col++ {
			if visited[row][col] {
				continue
			}

			if grid[row][col] == '1' {
				islandCount++
			}
			visited = dfsVisit(grid, visited, row, col)
		}
	}

	return islandCount
}

func makeVisitedGrid(rowCount, colCount int) [][]bool {
	visitedGrid := make([][]bool, rowCount)
	for i := 0; i < rowCount; i++ {
		visitedGrid[i] = make([]bool, colCount)
	}

	return visitedGrid
}

func dfsVisit(grid [][]byte, visited [][]bool, row, col int) [][]bool {
	if row >= len(grid) || col >= len(grid[0]) || row < 0 || col < 0 {
		return visited
	}

	if visited[row][col] {
		return visited
	}

	visited[row][col] = true
	if grid[row][col] == '1' {
		visited = dfsVisit(grid, visited, row+1, col)
		visited = dfsVisit(grid, visited, row, col+1)
		visited = dfsVisit(grid, visited, row-1, col)
		visited = dfsVisit(grid, visited, row, col-1)
	}
	return visited
}

func main() {
	fmt.Println(numIslands([][]byte{{'1', '1', '1'}, {'0', '1', '0'}, {'1', '1', '1'}}))
	fmt.Println(numIslands([][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}))
	fmt.Println(numIslands([][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}}))
}
