package problems

import (
	"fmt"
)

func Mapper(problem_number int) (func(input int) int, error) {
	switch problem_number {
	case 1:
		return sum_of_multiples, nil
	}
	return nil, fmt.Errorf("problem-%d not solved", problem_number)
}
