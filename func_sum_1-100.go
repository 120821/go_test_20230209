package main

import "fmt"

func main() {
  sum := 0
  for i := 1; i <= 100; i++ {
    fmt.Println("i is ", i)
    sum = sum + i
  }
  fmt.Printf( "sum: ", sum)
}

