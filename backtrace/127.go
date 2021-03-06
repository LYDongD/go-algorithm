package main

import "fmt"

var minCount int

func ladderLength(beginWord string, endWord string, wordList []string) int {
	if len(beginWord) != len(endWord) {
		return 0
	}

	if !endWordExist(endWord, wordList) {
		return 0
	}

	minCount = int(^uint(0) >> 1)
	backtrace(beginWord, endWord, wordList, 1)
	if minCount == int(^uint(0)>>1) {
		minCount = 0
	}
	return minCount
}

func endWordExist(endWord string, wordList []string) bool {
	for _, word := range wordList {
		if word == endWord {
			return true
		}
	}

	return false
}

func oneLetterDiff(left, right string) bool {
	differ := 0
	for i := 0; i < len(left); i++ {
		if left[i] != right[i] {
			differ++
		}
	}

	return differ == 1
}

func backtrace(currentWord, endWord string, wordList []string, count int) {

	if currentWord == endWord {
		if count < minCount {
			minCount = count
		}
		return
	}

	if len(wordList) == 0 {
		return
	}

	fmt.Println(currentWord, wordList)

	for i := 0; i < len(wordList); i++ {
		word := wordList[i]
		if !oneLetterDiff(word, currentWord) {
			continue
		}

		tmpList := make([]string, len(wordList))
		copy(tmpList, wordList)
		tmpList = append(tmpList[:i], tmpList[i+1:]...)
		backtrace(word, endWord, tmpList, count+1)
	}

	return
}

func main() {
	fmt.Println(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
	fmt.Println(ladderLength("qa", "sq", []string{"si","go","se","cm","so","ph","mt","db","mb","sb","kr","ln","tm","le","av","sm","ar","ci","ca","br","ti","ba","to","ra","fa","yo","ow","sn","ya","cr","po","fe","ho","ma","re","or","rn","au","ur","rh","sr","tc","lt","lo","as","fr","nb","yb","if","pb","ge","th","pm","rb","sh","co","ga","li","ha","hz","no","bi","di","hi","qa","pi","os","uh","wm","an","me","mo","na","la","st","er","sc","ne","mn","mi","am","ex","pt","io","be","fm","ta","tb","ni","mr","pa","he","lr","sq","ye"}))
}
