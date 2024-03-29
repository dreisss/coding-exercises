package main

func fib(n int) int {
	if n <= 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

// tests
func main() {
	println(fib(1))
	println(fib(2))
	println(fib(3))
	println(fib(4))
	println(fib(5))
	println(fib(6))
	println(fib(7))
	println(fib(8))
	println(fib(9))
	println(fib(10))
}
