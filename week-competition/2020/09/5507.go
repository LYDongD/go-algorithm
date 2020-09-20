package main

import "fmt"

func modifyString(s string) string {
	if len(s) == 0 {
		return s
	}

	for i, charRune := range s {
		char := byte(charRune)
		if char == '?' {
			neighbour := make(map[byte]int)
			if i-1 >= 0 && s[i-1] != '?' {
				neighbour[s[i-1]]++
			}

			if i+1 <= len(s)-1 && s[i+1] != '?' {
				neighbour[s[i+1]]++
			}

			for t := 0; t <= 25; t++ {
				toReplace := byte('a' + t)
				if neighbour[toReplace] == 0 {
					s = replace(s, i, toReplace)
					break
				}
			}
		}
	}

	return s
}

func replace(s string, index int, r byte) string {
	out := []byte(s)
	out[index] = r
	return string(out)
}

func main() {
	fmt.Println(modifyString("?zs"))
	fmt.Println(modifyString("ubv?w"))
	fmt.Println(modifyString("j?qg??b"))
	fmt.Println(modifyString("??yw?ipkj?"))
}
