package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
)

// const FILE_INPUT = "input.txt"
const FILE_INPUT = "/Users/jesse/Projects/Golang/learn-go/advent-of-code/day1/input.txt"

func main() {
	// puzzle1()
	var puzzle int
	puzzle, _ = strconv.Atoi(os.Args[1])
	switch puzzle {
	case 1:
		puzzle1()
	case 2:
		puzzle2()
	case 3:
		// NOTE: This section just tests out array slices in Go
		var arr [3]int
		arrayMuteTest(arr[:])
		for i := 0; i < len(arr); i++ {
			fmt.Println(arr[i])
		}
	default:
		fmt.Println("Invalid puzzle number")
	}
}

func arrayMuteTest(array []int) {
	array[0] = 1
	array[1] = 2
	array[3] = -1
}

func printList(myList *list.List) {
	for e := myList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func panicOnError(e error) {
	if e != nil {
		panic(e)
	}
}

// Extract the integer from a line in the input file
// @returns (val, isEmpty) = (int value of line, true if line was empty)
func parseLine(line string) (val int, isEmpty bool) {
	if len(line) == 0 {
		isEmpty = true
	} else {
		// WARNING: This ignores errors in strconv.Atoi
		isEmpty = false
		val, _ = strconv.Atoi(line)
	}

	return val, isEmpty
}

// math.Max() but for integers
func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func puzzle1() {
	var (
		file *os.File
		err  error
	)
	file, err = os.Open(FILE_INPUT)
	panicOnError(err)
	fmt.Println("File successfully opened: ", file.Name())
	var scanner *bufio.Scanner = bufio.NewScanner(file)

	var mostCalories, curCalorieSum int
	for scanner.Scan() {
		var (
			lineValue   int
			lineIsEmpty bool
		)

		lineValue, lineIsEmpty = parseLine(scanner.Text())
		if lineIsEmpty { // End of an elf's inventory
			// fmt.Printf("Comparing: %d VS %d\n", mostCalories, curCalorieSum)
			mostCalories = maxInt(mostCalories, curCalorieSum)
			curCalorieSum = 0
		} else {
			curCalorieSum += lineValue
		}

		// fmt.Printf("str=%s -- int=%d -- size=%d\n", text, textValue, len(text))
	}

	mostCalories = maxInt(mostCalories, curCalorieSum)
	fmt.Printf("Most calories: %d\n", mostCalories)
	file.Close()
}

func puzzle2() {
	// Find top 3 vals

	// WARN: Does not work! Arrays are static in Go. No(?) builtin list
	// Simple but inefficient:
	// 1) Sum all people in file. Save each to an array
	// 2) Sort the array in descending order
	// 3) Sum top 3 indices
	// TODO: Improved method, use 3 element array
	var (
		file    *os.File
		scanner *bufio.Scanner
		err     error
		// elfSums [3]int
	)

	// Load file into memory
	file, err = os.Open(FILE_INPUT)
	panicOnError(err)
	scanner = bufio.NewScanner(file)

	for scanner.Scan() {

	}
}
