package main

import "fmt"

type WordDictionary struct {
	recordMap map[string]int
}

func Constructor() WordDictionary {
	return WordDictionary{make(map[string]int)}
}

func (this *WordDictionary) AddWord(word string) {
	this.recordMap[word]++
}

func (this *WordDictionary) Search(word string) bool {
	if this.recordMap[word] > 0 {
		return true
	}

	if !isContain(word, '.') {
		return false
	}

	found := false
	for record, _ := range this.recordMap {
		if isMatch(record, word) {
			found = true
			break
		}
	}

	return found
}

func isMatch(wordA, wordB string) bool {
	if len(wordA) != len(wordB) {
		return false
	}

	for i := 0; i < len(wordA); i++ {
		if wordA[i] == '.' || wordB[i] == '.' || wordA[i] == wordB[i] {
			continue
		}

		return false
	}

	return true
}

func isContain(word string, char byte) bool {
	for _, c := range word {
		if byte(c) == char {
			return true
		}
	}

	return false
}

func main() {
	wordDic := Constructor()
	wordDic.AddWord("bad")
	wordDic.AddWord("dad")
	wordDic.AddWord("mad")
	fmt.Println(wordDic.Search("pad"))
	fmt.Println(wordDic.Search("bad"))
	fmt.Println(wordDic.Search(".ad"))
	fmt.Println(wordDic.Search("b.."))
}
