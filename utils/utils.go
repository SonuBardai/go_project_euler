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
