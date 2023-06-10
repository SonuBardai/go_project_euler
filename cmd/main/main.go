package main

import (
	"fmt"
	"os"
	"strconv"
)

func contains(items []string, item string) bool {
	for _, i := range items {
		if i == item {
			return true
		}
	}
	return false
}

func any(items, candidates []string) bool {
	for i := 0; i < len(candidates); i++ {
		if contains(items, candidates[i]) {
			return true
		}
	}
	return false
}

func main() {
	args := os.Args[1:]
	if any(args, []string{"--problem", "-p"}) {
		if problem_num, err := strconv.Atoi(args[1]); err == nil {
			fmt.Println("Problem: ", problem_num)
		}
	}
}
