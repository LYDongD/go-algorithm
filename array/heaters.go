package main

import "fmt"

func findRadius(houses []int, heaters []int) int {

	radius := -1
	for _, house := range houses {
		minRadius := minRadiusForHouse(house, heaters)
		if radius < 0 || radius < minRadius {
			radius = minRadius
		}
	}

	return radius
}

//find min warm radius for a house
func minRadiusForHouse(house int, heaters []int) int {

	minRadius := -1
	for _, heater := range heaters {
		distance := distance(house, heater)
		if minRadius < 0 || minRadius > distance {
			minRadius = distance
		}
	}

	return minRadius
}

func distance(a int, b int) int {
	result := a - b
	if result < 0 {
		result = -result
	}

	return result
}

func main() {
	houses := []int{1, 2, 3, 4}
	heaters := []int{1, 4}
	fmt.Println(findRadius(houses, heaters))
}
