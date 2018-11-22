//package main
// 
//import (
//    "fmt"
//    "strconv"
// )

func compress(chars []byte) int {

	if len(chars) == 0 {
		return 0
	}

	//point to position wait to be inserted and return as length in the end
	cursor := 0

	anchor := 0

	for index, char := range chars {
		if index == len(chars)-1 || chars[index+1] != char {
			chars[cursor] = chars[anchor]
			cursor++

			if index > anchor {
				for _, c := range strconv.Itoa(index - anchor + 1) {
					chars[cursor] = byte(c)
					cursor++
				}
			}

			anchor = index + 1
		}
	}

	return cursor
}

//func main() {
//
//	//chars := []byte{'a', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c'}
//	//chars := []byte{'a', 'b'}
//	chars := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
//	length := compress(chars)
//	fmt.Println(length)
//	fmt.Println(chars[:length])
//
//}
