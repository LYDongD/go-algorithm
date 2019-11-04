package main

import "fmt"

func findItinerary(tickets [][]string) []string {
	result := []string{}
	if len(tickets) == 0 {
		return result
	}

	dic := make(map[string][]string)
	for _, ticket := range tickets {
		from, to := ticket[0], ticket[1]
		if dic[from] == nil {
			dic[from] = []string{to}
		} else {
			toArr := dic[from]
			insertIndex := -1
			for i := 0; i < len(toArr); i++ {
				if toArr[i] > to {
					insertIndex = i
				}
			}
			if insertIndex >= 0 {
				rear := append([]string{}, toArr[insertIndex:]...)
				toArr = append(toArr[0:insertIndex], to)
				toArr = append(toArr, rear...)
			} else {
				toArr = append(toArr, to)
			}
			dic[from] = toArr

		}
	}

	nextFrom := "JFK"
	result = append(result, nextFrom)
	count := 0
	for count < len(tickets)+1 {
		toArr := dic[nextFrom]
		if len(toArr) > 0 {
			result = append(result, toArr[0])
			dic[nextFrom] = toArr[1:]
			nextFrom = toArr[0]
		}

		count++
	}

	return result
}

func main() {
	//fmt.Println(findItinerary([][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}}))
	fmt.Println(findItinerary([][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}}))
}
