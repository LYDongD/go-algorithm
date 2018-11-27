//package main
//
//import (
//	"fmt"
//	"math"
//)

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {

	//defens
	if buckets <= 1 || minutesToDie == 0 || minutesToTest == 0 {
		return 0
	}

	base := (minutesToTest / minutesToDie) + 1
	//pig coung as k
	k := 1

	for int(math.Pow(float64(base), float64(k))) < buckets {
		k++
	}

	return k

}

//func main() {
//	fmt.Println(poorPigs(25, 15, 60))
//}
