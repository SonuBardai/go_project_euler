package problems

import (
	"fmt"
)

// go: noinline
func Mapper(problemNumber int) (func(input int) int, error) {
	switch problemNumber {
	case 1:
		return sumOfMultiples, nil
	case 2:
		return fibSeries, nil
	case 3:
		return largestPrimeFactor, nil
	case 4:
		return largestPalindromeProduct, nil
	}
	return nil, fmt.Errorf("problem-%d not solved", problemNumber)
}
