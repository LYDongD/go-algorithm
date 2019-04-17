package main

import "fmt"

func orangesRotting(grid [][]int) int {
	if len(grid) == 0 {
		return -1
	}

	rotted := false
	count := 0
	for {
		current := makeNewGrid(grid)
		for i := 0; i < len(grid); i++ {
			row := grid[i]
			for j := 0; j < len(row); j++ {
				orange := row[j]
				if orange == 2 {
					if j > 0 && row[j-1] == 1 {
						current[i][j-1] = 2
						rotted = true
					}

					if j < len(row)-1 && row[j+1] == 1 {
						current[i][j+1] = 2
						rotted = true
					}

					if i > 0 && grid[i-1][j] == 1 {
						current[i-1][j] = 2
						rotted = true
					}

					if i < len(grid)-1 && grid[i+1][j] == 1 {
						current[i+1][j] = 2
						rotted = true
					}

				}

				if current[i][j] != 2 {
					current[i][j] = grid[i][j]
				}
			}
		}

		if !rotted {
			if hasFreshOrange(current) {
				return -1
			}

			return count
		}

		count++
		rotted = false
		grid = current
	}
}

func makeNewGrid(grid [][]int) [][]int {
	newGrid := make([][]int, len(grid))
	for i := 0; i < len(newGrid); i++ {
		newGrid[i] = make([]int, len(grid[0]))
	}
	return newGrid
}

func hasFreshOrange(grid [][]int) bool {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				return true
			}
		}
	}

	return false
}

func main() {
	fmt.Println(orangesRotting([][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}))
	fmt.Println(orangesRotting([][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}))
	fmt.Println(orangesRotting([][]int{{0, 2}}))
}
