package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
)

// const INPUT_FILE = "input.txt"
const INPUT_FILE = "/Users/jesse/Projects/Golang/learn-go/advent-of-code/day1/input.txt"

func main() {
	// puzzle1()
	var puzzle int
	puzzle, _ = strconv.Atoi(os.Args[1])
	switch puzzle {
	case 1:
		puzzle1()
	case 2:
		total := puzzle2()
		fmt.Println("Total of the top 3: ", total)
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
	file, err = os.Open(INPUT_FILE)
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

// Finds the sum of the top 3 values in INPUT_FILE
// METHOD:
//   - Define an array which holds the 3 largest value (topElfSums)
//   - Read the file *once*. For each elf in the input file:
//     -> Sum their "calories"
//     -> Use a modified Insertion Sort to insert each new value into `topElfSums`. O(k) where k= # of elves.
//   - Return the sum of ints in topElfSums => O(k) where k= # of elves
func puzzle2() int {
	var (
		file       *os.File
		scanner    *bufio.Scanner
		err        error
		topElfSums [3]int // NOTE: Invariant -- always sorted in descending order
		elfCurrSum int
	)

	// Load file into memory
	file, err = os.Open(INPUT_FILE)
	panicOnError(err)
	scanner = bufio.NewScanner(file)

	// Read through the file, summing up each elf & tracking the top 3
	for scanner.Scan() { // NOTE: O(N) -- only 1 file pass
		line := scanner.Text()

		// Empty lines separate each elf's inventory
		if len(line) == 0 {
			elfCountInsert(topElfSums[:], elfCurrSum)
			elfCurrSum = 0
		} else {
			lineVal, err := strconv.Atoi(line)
			panicOnError(err)
			elfCurrSum += lineVal
		}
	}

	// Don't forget the last elf! TODO: Incorporate into main loop
	if elfCurrSum != 0 {
		// NOTE: O(1) time; complexity does not grow with length of input file (number of total elves)
		elfCountInsert(topElfSums[:], elfCurrSum)
	}

	return topElfSums[0] + topElfSums[1] + topElfSums[2]
}

// Inserts newElf into topElves if it is a new top3 elf
// @WARN mutates topElves!
func elfCountInsert(topElves []int, newElf int) {
	// Insert at front
	// if new val > front:
	// push front down
	var idx int = 0
	var push int = newElf

	// TEST:
	// topElves [500, 200, 50]
	// newElf    100
	// -----
	// idx = 3
	// push = 100
	// len(topElves) = 3
	for idx < len(topElves) { // 3 < 3
		if push > topElves[idx] { // 100 > topElves[2]=50 -- TRUE
			tmp := topElves[idx] // tmp = 50
			topElves[idx] = push // topElves[2] =100
			push = tmp           // push = 50
		}
		idx++
	}
}
