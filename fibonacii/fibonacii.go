package internal

func fibonacii() func(x int) int {
	a, b := 0, 1
	return func(y int) int {
		result := a
		a, b = b, a+b
		return result
	}
}
