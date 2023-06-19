package palindrome

import (
	"math"
	"strconv"

	"github.com/SonuBardai/go_project_euler/utils"
)

func Niner(n int) int {
	var nines int
	for i := 0; i < n; i++ {
		nines += int(math.Pow(10, float64(i))) * 9
	}
	return nines
}

func IsPalindrome(n int) bool {
	str_n := strconv.Itoa(n)

	return str_n == utils.Reverse(str_n)

	// use pointers to find if string is palindrome
	// i := 0
	// j := len(str_n) - 1
	// for i <= j {
	// 	if str_n[i] == str_n[j] {
	// 		i += 1
	// 		j -= 1
	// 	} else {
	// 		return false
	// 	}
	// }
	// return true
}
