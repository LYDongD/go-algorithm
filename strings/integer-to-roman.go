package main

import "fmt"

func intToRoman(num int) string {

    if num == 0 {
	return ""
    }

    romanDic := map[int]string {
	1 : "I",
	5 : "V",
	10 : "X",
	50 : "L",
	100 : "C",
	500 : "D",
	1000 : "M",
    }

    base := 1
    result := ""
    for num > 0 {
	roman := ""
	current := num % 10
	if current == 4 {
	    roman = romanDic[base] + romanDic[5 * base]
	}else if current == 9 {
	    roman = romanDic[base] + romanDic[10 * base]
	}else {
	    if current >= 5 {
		roman = romanDic[5 * base]
		current -= 5
	    }

	    for i := 1; i <= current; i++ {
		roman =  roman + romanDic[base]
	    }
	}

	result = roman + result
	base = base * 10
	num = num / 10
    }

    return result
}


func main() {
    fmt.Println(intToRoman(3))
    fmt.Println(intToRoman(4))
    fmt.Println(intToRoman(9))
    fmt.Println(intToRoman(58))
    fmt.Println(intToRoman(1994))
    fmt.Println(intToRoman(2371))
}
