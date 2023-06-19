package problems

import (
	"fmt"
)

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
	case 5:
		return smallestMultiple, nil
	}
	return nil, fmt.Errorf("problem-%d not solved", problemNumber)
}
