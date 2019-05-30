package main

import "fmt"

var result []string

func restoreIpAddresses(s string) []string {
	result = []string{}

	backTrace(s, 1, "")

	return result
}

func backTrace(s string, selectedCount int, ipStr string) {
	if selectedCount == 4 {
		if len(s) > 3 || len(s) == 0 {
			return
		}

		if len(s) == 3 && s > "255" {
			return
		}

		if len(s) >= 2 && s[0] == '0' {
			return
		}

		result = append(result, ipStr+s)
		return
	}

	//select ip segment
	for i := 1; i <= 3; i++ {
		if len(s) < i {
			return
		}

		if i >= 2 && s[0] == '0' {
			return
		}

		if i == 3 && s[:3] > "255" {
			return
		}

		backTrace(s[i:], selectedCount+1, ipStr+s[:i]+".")
	}

}

func main() {
	fmt.Println(restoreIpAddresses("25525511135"))
	fmt.Println(restoreIpAddresses("0000"))
	fmt.Println(restoreIpAddresses("1111"))
}
