// NOTE:
// Exercise: Fibonacci closure
// Let's have some fun with functions.
// Implement a fibonacci function that returns a function (a closure)
// that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).package main
package exercises

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	var fib_0, fib_1 int = 0, 1
	counter := 0
	return func() int {
		var next int
		if counter == 0 {
			next = 0
		} else if counter == 1 {
			next = 1
		} else {
			next = fib_0 + fib_1
			fib_0 = fib_1
			fib_1 = next
		}
		counter += 1
		return next
	}
}

func Fibonacci() {
	f := fibonacci()
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(f())
	// }
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
