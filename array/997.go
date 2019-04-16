package main

import "fmt"

func findJudge(N int, trust [][]int) int {

	if N == 1 && len(trust) == 0 {
		return 1
	}

	if len(trust) < N-1 {
		return -1
	}

	//record candidate judge and it's trust count
	trustDic := make(map[int]int)

	for row := 0; row < len(trust); row++ {
		trustPair := trust[row]
		trustFrom := trustPair[0]
		trustTo := trustPair[1]
		trustDic[trustFrom] = -1
		if trustDic[trustTo] != -1 {
			trustDic[trustTo]++
		}
	}

	for candidate, count := range trustDic {
		if count > 0 && count == N-1 {
			return candidate
		}
	}

	return -1
}

func main() {
	fmt.Println(findJudge(4, [][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}, {4, 3}}))
	fmt.Println(findJudge(3, [][]int{{1, 2}, {2, 3}}))
	fmt.Println(findJudge(3, [][]int{{1, 3}, {2, 3}, {3, 1}}))
	fmt.Println(findJudge(3, [][]int{{1, 3}, {2, 3}}))
}
