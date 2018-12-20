package main

import "fmt"

func checkRecord(s string) bool {
	if len(s) == 0 {
		return false
	}

	ACount := 0
	LContinusCount := 0

	for _, char := range s {
		switch char {
		case 'A':
			ACount++
			if ACount > 1 {
				return false
			}
			LContinusCount = 0
		case 'L':
			LContinusCount++
			if LContinusCount > 2 {
				return false
			}
		default:
			LContinusCount = 0
		}

	}

	return true
}

func main() {
	fmt.Println(checkRecord("PPALLP"))
	fmt.Println(checkRecord("PPALLL"))
}
