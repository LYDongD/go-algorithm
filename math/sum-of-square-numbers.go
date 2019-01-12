package math

func judgeSquareSum(c int) bool {
    if c < 0 {
        return false
    }
    
    for a := 0;  a * a <= c; a++ {
        b := c - a * a
        if perfectSquare(b) {
            return true
        }
    }

    return false
}

func perfectSquare(n int) bool {
    
    start := 0
    end := n
    for start <= end  {
        middle := start + (end - start) / 2 
        m := middle * middle 
        if m == n {
            return true
        }else if m < n {
            start = middle + 1
        }else {
            end  = middle - 1
        }
    }

    return false
}
