package fib

//nolint:deadcode
func fibExponential(n int) int {
	if n <= 1 {
		return n
	} else {
		return fibExponential(n-1) + fibExponential(n-2)
	}
}

func FibOptimized(n int) int {
	a := 0
	b := 1
	count := 2
	for count <= n {
		c := a + b
		a = b
		b = c
		count += 1
	}
	return b
}
