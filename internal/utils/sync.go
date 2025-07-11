package utils

import "fmt"

func Calc(c, quit chan int) {
	x, y := 0, 1
	select {
	case c <- x:
		x, y = y, x+y

	case <-quit:
		fmt.Printf("quit")
		return
	}
}
