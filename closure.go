package main

func fibonacciRecursive(n int) int {
	if n <= 1 {
		return 1
	}
	return n + fibonacciRecursive(n-1)
}

//Solving Golang lack of TCO (Tail Call Optimization)
func fibonacci(n int) int {
	current, prev := 0, 1
	for i := 0; i < n; i++ {
		current, prev = current+prev, current
	}
	return current
}
