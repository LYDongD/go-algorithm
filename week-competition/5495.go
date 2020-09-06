package main

import "fmt"

func mostVisited(n int, rounds []int) []int {
	mostVistitedArr := []int{}
	if len(rounds) == 0 {
		return mostVistitedArr
	}

	countArr := make([]int, n)
	countArr[rounds[0]-1]++
	for i := 0; i < len(rounds)-1; i++ {
		if rounds[i]+1 <= rounds[i+1] {
			for k := rounds[i]; k <= rounds[i+1]-1; k++ {
				countArr[k]++
			}
		} else {
			for k := 0; k <= rounds[i+1]-1; k++ {
				countArr[k]++
			}

			for k := rounds[i]; k <= n-1; k++ {
				countArr[k]++
			}
		}
		//	fmt.Println("rount: ", i+1, "countArr", countArr)
	}

	//	fmt.Println(countArr)

	max := 0
	for i, count := range countArr {
		if count == max {
			mostVistitedArr = append(mostVistitedArr, i+1)
		} else if count > max {
			mostVistitedArr = []int{}
			mostVistitedArr = append(mostVistitedArr, i+1)
			max = count
		}
	}

	return mostVistitedArr
}

func main() {
	fmt.Println(mostVisited(4, []int{1, 3, 1, 2}))
	fmt.Println(mostVisited(2, []int{2, 1, 2, 1, 2, 1, 2, 1, 2}))
	fmt.Println(mostVisited(7, []int{1, 3, 5, 7}))
}
