package main

import "fmt"

type TrieNode struct {
	word     string
	children []*TrieNode
}

type WordDictionary struct {
	root *TrieNode
}

func Constructor() WordDictionary {

	rootNode := &TrieNode{
		word:     "",
		children: make([]*TrieNode, 26),
	}

	return WordDictionary{root: rootNode}
}

func (this *WordDictionary) AddWord(word string) {
	if len(word) == 0 {
		return
	}

	trieNode := this.root
	for _, char := range word {
		if trieNode.children[char-'a'] == nil {
			trieNode.children[char-'a'] = &TrieNode{children: make([]*TrieNode, 26)}
		}
		trieNode = trieNode.children[char-'a']
	}

	trieNode.word = word
}

func (this *WordDictionary) Search(word string) bool {
	if len(word) == 0 {
		return true
	}

	return dfsMatch(this.root, word, 0)
}

func dfsMatch(node *TrieNode, word string, k int) bool {
	if k == len(word) {
		return node.word != ""
	}

	if word[k] != '.' {
		return node.children[word[k]-'a'] != nil && dfsMatch(node.children[word[k]-'a'], word, k+1)
	}

	for i := 0; i < len(node.children); i++ {
		if node.children[i] != nil {
			isMatched := dfsMatch(node.children[i], word, k+1)
			if isMatched {
				return true
			}
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
