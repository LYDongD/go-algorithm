package main

import "fmt"

func convert(s string, numRows int) string {
	if len(s) <= 1 || numRows <= 1 {
		return s
	}

	//generate zigzag pattern img
	img := make([][]byte, numRows)

	for j, k := 0, 0; k < len(s); j++ {

		if j%(numRows-1) == 0 {
			for i := 0; i < numRows && k < len(s); i++ {
				img[i] = append(img[i], s[k])
				k++
			}
		} else {
			t := numRows - j%(numRows-1) - 1
			for i := 0; i < numRows && k < len(s); i++ {
				if i == t {
					img[i] = append(img[i], s[k])
					k++
				} else {
					img[i] = append(img[i], 0)
				}
			}
		}
	}

	//output
	byteResult := []byte{}
	for i := 0; i < len(img); i++ {
		for j := 0; j < len(img[i]); j++ {
			if img[i][j] != 0 {
				byteResult = append(byteResult, img[i][j])
			}
		}
	}

	return string(byteResult)

}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 3))
	fmt.Println(convert("PAYPALISHIRING", 4))
}
