package main

import (
	"fmt"
	"os"

	"github.com/jsteinberg4/learn-go/leetcode/problems"
)

var SOLUTIONS = map[string]func(){
	"test":             func() { fmt.Println("Test function") },
	"MergeSortedArray": problems.MergeSortedArray,
}

func main() {
	var (
		problemName string
		fn          func()
		hasFn       bool
	)

	// Check for correct function usage
	if len(os.Args) < 2 {
		fmt.Println("Error! Expected format: go run . [ProblemName]")
		panic("Incorrect # of cmdline arguments")
	}

	// Call the correct function
	problemName = os.Args[1]
	if fn, hasFn = SOLUTIONS[problemName]; hasFn {
		fmt.Println("Function found!")
		fn()
	} else {
		fmt.Println("Not found")
	}
}
