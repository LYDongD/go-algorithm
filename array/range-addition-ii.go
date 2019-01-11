package array

import "fmt"

func maxCount(m int, n int, ops [][]int) int {
	if len(ops) == 0 {
		return m * n
	}

	minRow, minCol := findIntersection(ops)

	return minRow * minCol
}

func findIntersection(ops [][]int) (int, int) {
	minRow := int(^uint(0) >> 1)
	minCol := int(^uint(0) >> 1)
	for _, op := range ops {
		if op[0] < minRow {
			minRow = op[0]
		}

		if op[1] < minCol {
			minCol = op[1]
		}
	}

	return minRow, minCol
}

func array() {
	operations := [][]int{{2, 2}, {3, 3}}
	minRow, minCol := findIntersection(operations)
	fmt.Println(minRow, minCol)
	fmt.Println(maxCount(3, 3, operations))

}
