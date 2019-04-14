package main

import "fmt"

func robotSim(commands []int, obstacles [][]int) int {
	if len(commands) == 0 {
		return 0
	}

	current := []int{0, 0}
	//0->N, 1->E, 2->S, 3->W
	direction := 0
	max := 0

	for _, command := range commands {
		switch command {
		case -1:
			if direction == 3 {
				direction = 0
			} else {
				direction++
			}

		case -2:
			if direction == 0 {
				direction = 3
			} else {
				direction--
			}
		default:
			for command > 0 {
				isObstacle := false
				current, isObstacle = moveRobot(current, direction, obstacles)
				dis := current[0]*current[0] + current[1]*current[1]
				if dis > max {
					max = dis
				}
				if isObstacle {
					break
				}
				command--
			}
		}
	}

	return max
}

func moveRobot(current []int, direction int, obstacles [][]int) ([]int, bool) {
	isObstacle := false
	switch direction {
	case 0:
		current[1] = current[1] + 1
		if detectObstacle(current, obstacles) {
			current[1] = current[1] - 1
			isObstacle = true
		}
	case 1:
		current[0] = current[0] + 1
		if detectObstacle(current, obstacles) {
			current[0] = current[0] - 1
			isObstacle = true
		}
	case 2:
		current[1] = current[1] - 1
		if detectObstacle(current, obstacles) {
			current[1] = current[1] + 1
			isObstacle = true
		}
	case 3:
		current[0] = current[0] - 1
		if detectObstacle(current, obstacles) {
			current[0] = current[0] + 1
			isObstacle = true
		}
	}

	return current, isObstacle
}

func detectObstacle(current []int, obstacles [][]int) bool {
	if len(obstacles) == 0 {
		return false
	}

	for _, obstacle := range obstacles {
		if current[0] == obstacle[0] && current[1] == obstacle[1] {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(robotSim([]int{4, -1, 3}, [][]int{}))
}
