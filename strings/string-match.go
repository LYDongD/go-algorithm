package main

import "fmt"

//return first index of match substring b to a
func stringMatchBM(a string, b string) int {

	n := len(a)
	m := len(b)
	if n < m {
		return -1
	}

	//configure bc map for pattern, record the last char index
	bcMap := make(map[byte]int)
	for index, char := range b {
		bcMap[byte(char)] = index
	}

	//generate suffix and prefix
	suffix, prefix := generateGSSuffixAndPrefix(b)

	//move pattern string until mathch
	i := 0
	for i <= n-m {

		//find bad character index
		j := findBCIndex(i, a, b)

		//no bad character, means match
		if j < 0 {
			return i
		}

		//cal distance
		bcX := calBCDistance(i, j, a, b, bcMap)
		gsX := calGSDistance(j, b, suffix, prefix)

		//distance to move
		x := max(bcX, gsX)
		i = i + x
	}

	return -1

}

func max(a int, b int) int {
	if a >= b {
		return a
	}

	return b
}

//todo
func generateGSSuffixAndPrefix(pattern string) ([]int, []bool) {

	m := len(pattern)

	suffix := make([]int, m-1)
	prefix := make([]bool, m-1)

	for index := range suffix {
		suffix[index] = -1
	}

	for i := 0; i < m-1; i++ {
		j := i
		k := 0
		for j >= 0 {
			if pattern[j] == pattern[m-k-1] {
				k++
				j--
			} else {
				break
			}

		}

		if k > 0 {
			suffix[k] = j + 1
		}

		if j < 0 {
			prefix[k] = true
		}
	}

	return suffix, prefix
}

func calBCDistance(start int, j int, main string, pattern string, bcMap map[byte]int) int {

	//find another bad char position, if no found, return -1
	anotherBCIndex := -1
	index, ok := bcMap[main[start+j]]
	if ok {
		anotherBCIndex = index
	}

	return j - anotherBCIndex
}

func calGSDistance(j int, pattern string, suffix []int, prefix []bool) int {

	m := len(pattern)

	//good suffix length
	k := m - j - 1

	//if there's another good suffix
	t := suffix[k]
	if t >= 0 {
		return j - t + 1
	}

	//if there's common prefix and suffix
	for s := j + 2; s < m; s++ {

		//find common prefix and suffix
		if prefix[m-s] {
			return s
		}
	}

	return -1
}

//find bad character from back to forward in pattern
func findBCIndex(start int, main string, pattern string) int {
	m := len(pattern)
	for j := m - 1; j >= 0; j-- {
		if main[start+j] != pattern[j] {
			return j
		}
	}

	return -1
}

func main() {
	fmt.Println("adccczytyhh", "tyh", stringMatchBM("adccczytyhh", "typ"))
}
