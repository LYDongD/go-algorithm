package main

import "fmt"

type LevelWord struct {
	word  string
	level int
}

type Queue struct {
	levelWords []*LevelWord
}

func (self *Queue) isEmpty() bool {
	return len(self.levelWords) == 0
}

func (self *Queue) offer(levelWord *LevelWord) {
	self.levelWords = append(self.levelWords, levelWord)
}

func (self *Queue) take() *LevelWord {
	if self.isEmpty() {
		panic("queue is empty")
	}

	first := self.levelWords[0]
	self.levelWords = self.levelWords[1:]
	return first
}

func ladderLength(beginWord string, endWord string, wordList []string) int {

	adjacentDic := generateAdjacentDic(beginWord, endWord, wordList)
	visited := make(map[string]int)
	queue := &Queue{[]*LevelWord{}}
	queue.offer(&LevelWord{beginWord, 1})
	for !queue.isEmpty() {
		current := queue.take()
		if current.word == endWord {
			return current.level
		}

		adjacentList := adjacentDic[current.word]
		for _, word := range adjacentList {
			if visited[word] == 0 {
				visited[word]++
				queue.offer(&LevelWord{word, current.level + 1})
			}
		}
	}

	return 0
}

func generateAdjacentDic(beginWord, endWord string, wordList []string) map[string][]string {
	//handle begin word
	dic := make(map[string][]string)
	dic[beginWord] = []string{}
	for _, word := range wordList {
		if differ(beginWord, word) == 1 {
			dic[beginWord] = append(dic[beginWord], word)
		}
	}

	//handle word in wordlist
	for _, word := range wordList {
		if word == endWord {
			continue
		}

		dic[word] = []string{}
		for _, candidate := range wordList {
			if differ(word, candidate) == 1 {
				dic[word] = append(dic[word], candidate)
			}
		}
	}

	return dic
}

func differ(left, right string) int {
	delta := 0
	for i := 0; i < len(left); i++ {
		if left[i] != right[i] {
			delta++
		}
	}
	return delta
}

func main() {
	fmt.Println(ladderLength("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
	fmt.Println(ladderLength("qa", "sq", []string{"si", "go", "se", "cm", "so", "ph", "mt", "db", "mb", "sb", "kr", "ln", "tm", "le", "av", "sm", "ar", "ci", "ca", "br", "ti", "ba", "to", "ra", "fa", "yo", "ow", "sn", "ya", "cr", "po", "fe", "ho", "ma", "re", "or", "rn", "au", "ur", "rh", "sr", "tc", "lt", "lo", "as", "fr", "nb", "yb", "if", "pb", "ge", "th", "pm", "rb", "sh", "co", "ga", "li", "ha", "hz", "no", "bi", "di", "hi", "qa", "pi", "os", "uh", "wm", "an", "me", "mo", "na", "la", "st", "er", "sc", "ne", "mn", "mi", "am", "ex", "pt", "io", "be", "fm", "ta", "tb", "ni", "mr", "pa", "he", "lr", "sq", "ye"}))
}
