package fibo

type FiboFn func(int) int

var fibMem = Memoize(fib)

func FibMemoized(n int) int {
	return fibMem(n)
}

func Memoize(f FiboFn) FiboFn {
	cache := make(map[int]int)
	return func(key int) int {
		if val, found := cache[key]; found {
			return val
		}
		temp := f(key)
		cache[key] = temp
		return temp
	}
}

func fib(x int) int {
	if x == 0 {
		return 0
	} else if x <= 2 {
		return 1
	} else {
		return fib(x-2) + fib(x-1)
	}
}
