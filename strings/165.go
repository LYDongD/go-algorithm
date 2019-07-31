package main

import "fmt"
import "strings"
import "strconv"

func compareVersion(version1 string, version2 string) int {
	version1Array := strings.Split(version1, ".")
	version2Array := strings.Split(version2, ".")

	size := max(len(version1Array), len(version2Array))

	for i := 0; i < size; i++ {

		versionNum1, versionNum2 := 0, 0
		if i < len(version1Array) {
			versionNum1, _ = strconv.Atoi(version1Array[i])
		}

		if i < len(version2Array) {
			versionNum2, _ = strconv.Atoi(version2Array[i])
		}

		if versionNum1 > versionNum2 {
			return 1
		} else if versionNum1 < versionNum2 {
			return -1
		}
	}

	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	fmt.Println(compareVersion("0.1", "1.1"))
	fmt.Println(compareVersion("1.0.1", "1"))
	fmt.Println(compareVersion("7.5.2.4", "7.5.3"))
	fmt.Println(compareVersion("1.01", "1.001"))
	fmt.Println(compareVersion("1.0", "1.0.0"))
}
