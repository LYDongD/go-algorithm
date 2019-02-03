package main

import "fmt"
import "strconv"

func calPoints(ops []string) int {
	if len(ops) == 0 {
		return 0
	}

	//generate round score
	var pointsArr []int
	for _, op := range ops {
		n := len(pointsArr)
		switch op {
		case "C":
			pointsArr = pointsArr[:n-1]
		case "+":
			pointsArr = append(pointsArr, pointsArr[n-2]+pointsArr[n-1])
		case "D":
			pointsArr = append(pointsArr, 2*pointsArr[n-1])
		default:
			num, err := strconv.Atoi(op)
			if err == nil {
				pointsArr = append(pointsArr, num)
			}
		}
	}

	//sum
	sum := 0
	for _, point := range pointsArr {
		sum += point
	}

	return sum
}

func main() {
	fmt.Println(calPoints([]string{"5", "2", "C", "D", "+"}))
	fmt.Println(calPoints([]string{"5", "-2", "4", "C", "D", "9", "+", "+"}))
}
