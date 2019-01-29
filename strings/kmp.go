package main

import "fmt"

func kmpMatch(a string, b string) int {
	if len(a) < len(b) {
		return -1
	}

	n := len(a)
	m := len(b)

	//j是指向模式串的指针，改变j的值使得模式串向后移动
	j := 0

	//计算失效函数，获取一个移位数组，根据坏字符的位置，获取好前缀的公共前缀下标，j移动到该下标后面
	next := getNext(b)

	//i是主串的指针，总是指向最后一个坏字符
	for i := 0; i < n; i++ {

		//如果遇到坏字符，需要移动模式串，根据好前缀原则
		for j > 0 && a[i] != b[j] {
			j = next[j-1] + 1
		}

		if a[i] == b[j] {
			j++
		}

		//模式串完全匹配
		if j == m {
			return i - m + 1
		}
	}

	return -1
}

//数学归纳法
func getNext(pattern string) []int {
	m := len(pattern)
	next := make([]int, m)
	next[0] = -1

	//k指向公共前缀的尾部
	k := -1
	for i := 1; i < m; i++ {

		for k >= 0 && pattern[i] != pattern[k+1] {
			k = next[k]
		}

		if pattern[i] == pattern[k+1] {
			k++
		}

		next[i] = k
	}

	return next
}

func main() {
	fmt.Println(kmpMatch("abcdss", "acd"))
}
