package main

import "fmt"
import "strconv"

func isAdditiveNumber(num string) bool {
    if len(num) < 3 {
	return false
    }

    return isAddictive(num, -1, -1, 0)
}

func isAddictive(leftNumStr string, sum, lastNum, depth int) bool {
    if len(leftNumStr) == 0 && depth > 2 {
	return true
    }


    for i := 0; i < len(leftNumStr); i++ {
	tryNum, _ := strconv.Atoi(leftNumStr[0:i+1])
	if sum >= 0 && tryNum != sum {
	    if tryNum == 0 {
		break
	    }else {
		continue
	    }
	}

	//fmt.Println("select try num:", tryNum, "sum", sum, "depth", depth)
	 result := isAddictive(leftNumStr[i+1:], plus(tryNum, lastNum), tryNum, depth + 1)
	 if result == true {
	    return true
	 }

	///no need to check 02 like num
	 if tryNum == 0 {
	    break
	 }
    }

    return false
}

func plus(tryNum, lastNum int) int {
    //first num handle, skip sum
    if lastNum == -1 {
	return -1
    }

    return tryNum + lastNum
}

func main(){
    fmt.Println(isAdditiveNumber("1203"))
    fmt.Println(isAdditiveNumber("1023"))
    fmt.Println(isAdditiveNumber("111"))
    fmt.Println(isAdditiveNumber("112358"))
    fmt.Println(isAdditiveNumber("199100199"))
}
