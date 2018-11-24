package main

import (
	"fmt"
	"math"
)

func numberOfBoomerangs(points [][]int) int {

	sum := 0
	for i, point1 := range points {
		distantCountMap := make(map[int]int)
		for j, point2 := range points {
			if i == j {
				continue
			}

			distant := distant(point1, point2)
			count := distantCountMap[distant]
			count++
			distantCountMap[distant] = count

		}

		for _, count := range distantCountMap {
			sum += count * (count - 1)
		}
	}

	return sum
}

//count distinct betweent two points
func distant(point1 []int, point2 []int) int {
	xDistinct := math.Pow(float64(point1[0]-point2[0]), float64(2))
	yDistinct := math.Pow(float64(point1[1]-point2[1]), float64(2))
	return int(xDistinct + yDistinct)
}

func main() {

	point1 := []int{1, 0}
	point2 := []int{0, 0}
	point3 := []int{2, 0}

	points := [][]int{point1, point2, point3}
	fmt.Println(points)
	fmt.Println(numberOfBoomerangs(points))

}
