package go_algorithm

func findNthDigit(n int) int {
	
	group := 1
	groupCount := 9
	groupStart := 1

	//find the start group
	for n > group * groupCount {
		n -= group * groupCount
		group++
		groupStart *= 10
		groupCount *= 10
	}

	len := group
	numStart := groupStart

	// find the start in group
	numStart += (n - 1) / len
	//offset begins from 0
	offset := (n - 1) % len

	 //find the digit 
	for i := 0; i < len - offset - 1; i++ {
		numStart /= 10
	}
	
	return numStart % 10
}
