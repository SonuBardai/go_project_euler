package problems

func sum_of_multiples(input int) int {
	var sum int
	for i := 0; i < input; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	return sum
}