package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/SonuBardai/go_project_euler/utils"
)

func main() {
	args := os.Args[1:]

	fmt.Println(args)

	if utils.Any(args, []string{"--problem", "-p"}) {
		if problem_num, err := strconv.Atoi(args[1]); err == nil {
			fmt.Println("Problem: ", problem_num)
		}
	}
}
