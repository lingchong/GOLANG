package utils

// 首字母大写表示导出函数
func Fibonacii() func(x int) int {
	a, b := 0, 1
	return func(y int) int {
		result := a
		a, b = b, a+b
		return result
	}
}
