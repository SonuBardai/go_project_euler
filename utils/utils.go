package utils

func contains(items []string, item string) bool {
	for _, i := range items {
		if i == item {
			return true
		}
	}
	return false
}

func Any(items, candidates []string) bool {
	for i := 0; i < len(candidates); i++ {
		if contains(items, candidates[i]) {
			return true
		}
	}
	return false
}

func Reverse(str string) string {
	r := []rune(str)
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
