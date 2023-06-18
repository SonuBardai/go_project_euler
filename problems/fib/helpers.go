package fib

// func _fibExponential(n int) int {
// 	if n <= 1 {
// 		return n
// 	} else {
// 		return _fibExponential(n-1) + _fibExponential(n-2)
// 	}
// }

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
