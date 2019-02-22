package main

import "fmt"

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	if maxChoosableInteger >= desiredTotal {
		return true
	}

	if maxChoosableInteger*(maxChoosableInteger+1)/2 < desiredTotal {
		return false
	}

	selectdResultCache := make(map[int]bool)
	return dfs(maxChoosableInteger, selectdResultCache, 0, desiredTotal)

}

func dfs(maxChoosableInteger int, selectdResultCache map[int]bool, selected int, desiredTotal int) bool {
	if desiredTotal <= 0 {
		return false
	}

	//if selected result calculated, return cache
	cacheResult, ok := selectdResultCache[selected]
	if ok {
		return cacheResult
	}

	//get result by competor result
	for i := 1; i <= maxChoosableInteger; i++ {
		//if current num selected, continue
		if selected&(1<<uint(i)) > 0 {
			continue
		}

		if !dfs(maxChoosableInteger, selectdResultCache, selected|(1<<uint(i)), desiredTotal-i) {
			selectdResultCache[selected] = true
			return true
		}
	}

	selectdResultCache[selected] = false
	return false
}
