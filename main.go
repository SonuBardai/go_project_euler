package main

import (
	"flag"
	"fmt"

	"github.com/SonuBardai/go_project_euler/problems"
)

func main() {
	var input int
	var problem int
	flag.IntVar(&problem, "problem", 1, "Problem number")
	flag.IntVar(&input, "input", 0, "Value of input to the problem")
	flag.Parse()
	Solution, err := problems.Mapper(problem)
	if err != nil {
		fmt.Println(err)
	} else {

		solution := Solution(input)
		fmt.Printf("Solution: %d\n", solution)
	}
}
