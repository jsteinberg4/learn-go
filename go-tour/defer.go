package main

import "fmt"

func main()  {

  defer fmt.Println("Goodbye :(")
  defer fmt.Println("Pretend you're doing some stuff....")
  defer fmt.Println("Hello! This is one of my first golang programs :)")
}
