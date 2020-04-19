package main

import "fmt"

func isNumber(s string) bool {

    s = trim(s)
    if len(s) == 0 {
	return false
    }

    numCount,dotCount, eCount := 0, 0, 0
    for index, char := range s {
	if isNum(char) {
	    numCount++
	}else if isDot(char) {
	    if dotCount > 0 || eCount > 0 {
		return false
	    }

	    if index > 0 && isExponent(rune(s[index-1])) {
		return false
	    }

	    dotCount++
	}else if isSign(char)  {
	    if index != 0 {
		if !isExponent(rune(s[index-1])) {
		    return false
		}

		if index == len(s) -1 || !isNum(rune(s[index+1])) {
		    return false
		}
	    }

	}else if isExponent(char)  {
	    if eCount > 0 || numCount == 0 {
		return false
	    }

	    if index  == 0 || index == len(s) - 1 {
		return false
	    }

	    eCount++
	}else {
	    return false
	}
    }

    return numCount > 0
}

func trim(s string) string {
    preSpace, postSpace := 0, 0
    for i := 0; i < len(s); i++ {
        if s[i] == ' ' {
            preSpace++
        }else {
            break
        }
    }

    if preSpace == len(s) {
        return ""
    }

    for i := len(s) - 1; i >= 0; i-- {
        if s[i] == ' ' {
            postSpace++
        }else {
            break
        }
    }

    return s[preSpace:len(s) - postSpace]
}

func isExponent(char rune) bool {
    return char == 'e' || char == 'E'
}

func isNum(char rune) bool {
    return char >= '0' && char <= '9'
}

func isSign(char rune) bool {
    return char == '+' || char == '-'
}

func isDot(char rune) bool {
    return char == '.'
}

func main() {
    fmt.Println(isNumber("-.7e+0435"))
    fmt.Println(isNumber("-1."))
    fmt.Println(isNumber("005047e+6"))
    fmt.Println(isNumber("+0e-"))
    fmt.Println(isNumber("1e."))
    fmt.Println(isNumber(" "))
    fmt.Println(isNumber(". 1"))
    fmt.Println(isNumber(".1"))
    fmt.Println(isNumber("3."))
    fmt.Println(isNumber("1"))
    fmt.Println(isNumber("+100"))
    fmt.Println(isNumber("5e2"))
    fmt.Println(isNumber("-123"))
    fmt.Println(isNumber("3.1416"))
    fmt.Println(isNumber("0123"))
    fmt.Println(isNumber("-1E-16"))
    fmt.Println(isNumber("12e"))
    fmt.Println(isNumber("1a3.14"))
    fmt.Println(isNumber("1.2.3"))
    fmt.Println(isNumber("+-5"))
    fmt.Println(isNumber("12e+5.4"))
}
