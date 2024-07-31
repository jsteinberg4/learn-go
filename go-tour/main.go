package main

import (
	"flag"
	"fmt"

	"github.com/jsteinberg4/learn-go/go-tour/exercises"
)

func main() {
	exercise := flag.String("exercise", "none", "exercise name")
	flag.Parse()

	switch *exercise {
	case "tree":
		exercises.TreeMain()
	default:
		fmt.Println("Unrecognized exercise name " + *exercise)
	}
}
