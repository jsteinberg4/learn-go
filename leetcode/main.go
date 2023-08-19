package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/jsteinberg4/learn-go/leetcode/problems"
)

var SOLUTIONS = map[string]func(){
	"test":             func() { fmt.Println("Test function") },
	"MergeSortedArray": problems.MergeSortedArray,
	"RemoveElement":    problems.RemoveElement,
}

func main() {
	var (
		problemName string
		fn          func()
		hasFn       bool
	)

	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	slog.SetDefault(slog.New(handler))

	// Check for correct function usage
	if len(os.Args) < 2 {
		fmt.Println("Error! Expected format: go run . [ProblemName]")
		panic("Incorrect # of cmdline arguments")
	}

	// Call the correct function
	problemName = os.Args[1]
	if fn, hasFn = SOLUTIONS[problemName]; hasFn {
		slog.Info(fmt.Sprintf("Function found: %v", problemName))
		fn()
	} else {
		fmt.Println("Not found")
	}
}
