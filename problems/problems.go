package problems

import (
	"fmt"
)

// go: noinline
func Mapper(problemNumber int) (func(input int) int, error) {
	switch problemNumber {
	case 1:
		return sumOfMultiples, nil
	}
	return nil, fmt.Errorf("problem-%d not solved", problemNumber)
}
