package main

import "fmt"

type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	val       rune
	children  []*TrieNode
	isEndChar bool
}

func Constructor() Trie {
	return Trie{&TrieNode{children: make([]*TrieNode, 26)}}
}

func (this *Trie) Insert(word string) {
	if len(word) == 0 {
		return
	}

	trieNode := this.root
	for _, char := range word {
		if trieNode.children[char-'a'] == nil {
			trieNode.children[char-'a'] = &TrieNode{char, make([]*TrieNode, 26), false}
		}
		trieNode = trieNode.children[char-'a']
	}

	trieNode.isEndChar = true
}

func (this *Trie) Search(word string) bool {
	if len(word) == 0 {
		return false
	}

	trieNode := this.root
	for _, char := range word {
		if trieNode.children[char-'a'] == nil {
			return false
		}
		trieNode = trieNode.children[char-'a']
	}

	return trieNode.isEndChar
}

func (this *Trie) StartsWith(prefix string) bool {
	if len(prefix) == 0 {
		return true
	}

	trieNode := this.root
	for _, char := range prefix {
		if trieNode.children[char-'a'] == nil {
			return false
		}
		trieNode = trieNode.children[char-'a']
	}
	return true
}

func main() {
	fmt.Println("vim-go")
}
