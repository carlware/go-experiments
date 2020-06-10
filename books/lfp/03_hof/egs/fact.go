package egs

import "fmt"

func factorial(x int, next func(int)) {
	if x == 0 {
		next(1)
	} else {
		factorial(x-1, func(y int) {
			next(x * y)
		})
	}
}

func TestFactorial() {
	factorial(4, func(result int) {
		fmt.Println(result)
	})
}
