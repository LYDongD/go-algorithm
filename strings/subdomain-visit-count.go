package main

import "fmt"
import "strconv"

func subdomainVisits(cpdomains []string) []string {
	if len(cpdomains) == 0 {
		return cpdomains
	}

	wordCountMap := wordCount(cpdomains)

	var result []string
	for word, count := range wordCountMap {
		countWord := strconv.Itoa(count) + " " + word
		result = append(result, countWord)
	}

	return result
}

func wordCount(cpdomains []string) map[string]int {
	wordCountMap := make(map[string]int)

	for _, domainCount := range cpdomains {
		spaceIndex := findSpace(domainCount)
		count, err := strconv.Atoi(domainCount[:spaceIndex])
		if err != nil {
			return nil
		}

		domain := domainCount[spaceIndex+1:]

		for i := len(domain) - 1; i >= 0; i-- {
			if domain[i] == byte('.') {
				wordCountMap[domain[i+1:]] += count
			}
		}

		wordCountMap[domain] += count
	}

	return wordCountMap
}

func findSpace(word string) int {
	for i := 0; i < len(word); i++ {
		if word[i] == byte(' ') {
			return i
		}
	}

	return -1
}

func main() {
	fmt.Println(subdomainVisits([]string{"9001 discuss.leetcode.com"}))
	fmt.Println(subdomainVisits([]string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}))
}
