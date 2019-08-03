package main

import "fmt"
import "strconv"

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}

	basePosDic := make(map[int]int)
	result := ""
	result += getSign(numerator, denominator)
	numerator, denominator = abs(numerator), abs(denominator)
	result += strconv.Itoa(numerator / denominator)
	base := numerator % denominator
	if base == 0 {
		return result
	}

	result += "."
	for base != 0 {
		if basePosDic[base] > 0 {
			position := basePosDic[base]
			result = result[:position] + "(" + result[position:]
			result += ")"
			return result
		}

		decimalNum := base * 10 / denominator
		result += strconv.Itoa(decimalNum)
		basePosDic[base] = len(result) - 1
		base = base * 10 % denominator
	}

	return result

}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func getSign(a, b int) string {
	if a*b < 0 {
		return "-"
	}

	return ""
}

func main() {
	fmt.Println(fractionToDecimal(1, 2))
	fmt.Println(fractionToDecimal(2, 1))
	fmt.Println(fractionToDecimal(2, 3))
	fmt.Println(fractionToDecimal(999, 13))
	fmt.Println(fractionToDecimal(-50, 8))
}
