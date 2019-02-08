package main

import "fmt"

//trie数节点
type Node struct {
	word  string
	chars [26]*Node
}

type Trie struct {
	root *Node
}

//插入单词，前缀优先插入节点，不重复插入
func (this *Trie) insertChar(word string) {
	cursor := this.root
	for _, char := range word {
		if cursor.chars[char-'a'] == nil {
			cursor.chars[char-'a'] = &Node{}
		}
		cursor = cursor.chars[char-'a']
	}

	cursor.word = word
}

//BFS search use queue
func (this *Trie) findLongest(words []string) string {
	queue := []*Node{this.root}
	result := ""
	for len(queue) > 0 {
		//dequeue
		cursor := queue[0]
		queue = queue[1:]

		//backwards, make sure smallest order
		for i := len(cursor.chars) - 1; i >= 0; i-- {
			node := cursor.chars[i]
			if node != nil && len(node.word) > 0 {
				result = node.word
				//enqueue
				queue = append(queue, node)
			}
		}

	}

	return result
}

func longestWord(words []string) string {
	trie := &Trie{&Node{}}
	for _, word := range words {
		trie.insertChar(word)
	}
	return trie.findLongest(words)
}

func main() {
	fmt.Println(longestWord([]string{"w", "wo", "wor", "worl", "world"}))
	fmt.Println(longestWord([]string{"a", "banana", "app", "appl", "ap", "apply", "apple"}))
}
