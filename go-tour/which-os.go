package main

import (
  "fmt"
  "runtime"
)

func main()  {
  fmt.Print("Go runs on: ")
  switch os := runtime.GOOS; os {
  case "darwin":
    fmt.Print("MacOS.\n")
  default:
    fmt.Printf("%s.\n", os)
    
  }
}
