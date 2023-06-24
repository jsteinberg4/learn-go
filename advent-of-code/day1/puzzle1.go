package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// const FILE_INPUT = "input.txt"
const FILE_INPUT = "/Users/jesse/Projects/Golang/learn-go/advent-of-code/day1/input.txt"

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

func main() {
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
