package array

import (
	"fmt"
	"math"
)

func constructRectangle(area int) []int {

	if area == 0 {
		return nil
	}

	result := make([]int, 2)
	l := int(math.Sqrt(float64(area)))
	w := l
	for l*w != area {
		if l*w > area {
			w--
		} else {
			l++
		}
	}

	result[0] = l
	result[1] = w

	return result

}

func array() {

	fmt.Println(constructRectangle(4))
	fmt.Println(constructRectangle(30))
	fmt.Println(constructRectangle(10000000))
}
