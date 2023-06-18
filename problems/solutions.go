package problems

import (
	"math"

	"github.com/SonuBardai/go_project_euler/problems/fib"
)

func sumOfMultiples(input int) int {
	var sum int
	for i := 0; i < input; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}

func fibSeries(input int) int {
	var series []int
	i := 0
	for {
		next := fib.FibOptimized(i)
		if next >= input {
			break
		}
		series = append(series, next)
		i += 1
	}
	var sum int
	// var even []int
	for i := 0; i < len(series); i++ {
		if series[i]%2 == 0 {
			sum += series[i]
			// even = append(even, series[i])
		}
	}
	// fmt.Println("Fibonacci series", series)
	// fmt.Printf("Even fibonacci numbers lower than %d: %v \n", input, even)
	return sum
}

func largestPrimeFactor(input int) int {
	maxPrime := 2
	n := input
	for n%2 == 0 {
		n /= 2
	}
	for i := 3; i <= int(math.Sqrt(float64(n))); i++ {
		if i%2 == 0 {
			continue
		}
		for n%i == 0 {
			maxPrime = i
			n /= i
		}
	}
	if n > 2 {
		maxPrime = n
	}
	return maxPrime
}
