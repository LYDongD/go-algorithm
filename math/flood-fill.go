package main

import "fmt"

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {

	if len(image) == 0 {
		return image
	}

	originColor := image[sr][sc]
	return fillColor(image, sr, sc, originColor, newColor)
}

func fillColor(image [][]int, sr int, sc int, originColor int, newColor int) [][]int {

	if !checkBounds(image, sr, sc) {
		return image
	}

	color := image[sr][sc]
	if color != originColor || color == newColor {
		return image
	}

	image[sr][sc] = newColor

	image = fillColor(image, sr-1, sc, originColor, newColor)
	image = fillColor(image, sr+1, sc, originColor, newColor)
	image = fillColor(image, sr, sc-1, originColor, newColor)
	image = fillColor(image, sr, sc+1, originColor, newColor)

	return image
}

func checkBounds(image [][]int, sr int, sc int) bool {
	return sr >= 0 && sr < len(image) && sc >= 0 && sc < len(image[0])
}

func main() {
	image := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	fmt.Println(floodFill(image, 1, 1, 2))
}
