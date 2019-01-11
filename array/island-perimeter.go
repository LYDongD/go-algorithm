package array

//
//import (
//	"fmt"
//)

func islandPerimeter(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	perimeter := 0
	for row := range grid {
		for column := range grid[row] {
			perimeter += sideLength(grid, row, column)
		}
	}

	return perimeter
}

func sideLength(grid [][]int, row int, column int) int {

	//water -> no side
	if grid[row][column] == 0 {
		return 0
	}

	//consider 4 sides
	sideLength := 0

	//left
	if column == 0 || grid[row][column-1] == 0 {
		sideLength++
	}

	//right
	if column == len(grid[row])-1 || grid[row][column+1] == 0 {
		sideLength++
	}

	//top
	if row == 0 || grid[row-1][column] == 0 {
		sideLength++
	}

	//bottom
	if row == len(grid)-1 || grid[row+1][column] == 0 {
		sideLength++
	}

	return sideLength
}

//func array() {
//	grid := [][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}
//	fmt.Println(islandPerimeter(grid))
//}
