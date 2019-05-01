package main

import "fmt"

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	if r0 > R || c0 > C || r0 < 0 || c0 < 0 {
		return nil
	}

	maxDis := 0
	disMap := make(map[int][][]int)
	//group by distance
	for r := 0; r < R; r++ {
		for c := 0; c < C; c++ {
			dis := calDis(r, c, r0, c0)
			if disMap[dis] == nil {
				disMap[dis] = [][]int{{r, c}}
			} else {
				disMap[dis] = append(disMap[dis], []int{r, c})
			}

			if dis > maxDis {
				maxDis = dis
			}
		}
	}

	//combine group in acend order
	var result [][]int
	for i := 0; i <= maxDis; i++ {
		group := disMap[i]
		for _, ele := range group {
			result = append(result, ele)
		}
	}

	return result
}

func calDis(R, C, r0, c0 int) int {
	return abs(R, r0) + abs(C, c0)
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}

	return b - a
}

func main() {
	fmt.Println(allCellsDistOrder(1, 2, 0, 0))
	fmt.Println(allCellsDistOrder(2, 2, 0, 1))
	fmt.Println(allCellsDistOrder(2, 3, 1, 2))
}
