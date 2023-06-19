package problems

import (
	"fmt"
	"math"

	"github.com/SonuBardai/go_project_euler/problems/divisors"
	"github.com/SonuBardai/go_project_euler/problems/fib"
	"github.com/SonuBardai/go_project_euler/problems/palindrome"
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

func largestPalindromeProduct(input int) int {
	if input > 8 {
		panic("That's gonna be a big number!")
	}
	for a := palindrome.Niner(input); a > palindrome.Niner(input-1); a-- {
		for b := a; b >= a-int(math.Pow10(input-1)); b-- {
			c := a * b
			if palindrome.IsPalindrome(c) {
				fmt.Println(a, b)
				return c
			}
		}
	}
	panic("Can't find palindrome!")
}

func smallestMultiple(input int) int {
	result := 1
	for i := 2; i <= input; i++ {
		result = divisors.Lcm(result, i)
	}
	return result
}
