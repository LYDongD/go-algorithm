package go_algorithm

//牛顿迭代求平方根
func Sqrt(x float64) float64 {
    z := 0.0
    for i := 0; i < 1000; i++ {
        z -= (z * z - x) / (2 * x)
    }

    return z

}
