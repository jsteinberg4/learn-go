package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/jsteinberg4/learn-go/leetcode/problems"
)

var (
	solutions = map[string]func(){
		"test":             func() { fmt.Println("Test function") },
		"MergeSortedArray": problems.MergeSortedArray,
		"RemoveElement":    problems.RemoveElement,
	}
	problemNameFlag *string
)

func init() {
	problemNameFlag = flag.String("p", "test", "Problem name")
}

func main() {
	// Parse CLI args
	flag.Parse()

	var (
		fn    func()
		hasFn bool
	)

	// Configure default logging
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	slog.SetDefault(slog.New(handler))

	// Call the correct function
	if fn, hasFn = solutions[*problemNameFlag]; hasFn {
		slog.Info(fmt.Sprintf("Function found: %s\n", *problemNameFlag))
		fn()
	} else {
		fmt.Printf("Function not found: %s\n", *problemNameFlag)
		os.Exit(1)
	}
}
