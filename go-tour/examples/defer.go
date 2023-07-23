package examples

import "fmt"

func DeferExample() {

	defer fmt.Println("Goodbye :(")
	defer fmt.Println("Pretend you're doing some stuff....")
	defer fmt.Println("Hello! This is one of my first golang programs :)")
}
