package main

import (
  "fmt"
)

func main() {
  if 7 % 2 == 0 {
    fmt.Println("7 is enve")
  } else {
    fmt.Println("7 is odd")
  }

  if 8 % 2 == 0 {
    fmt.Println("8 is enve")
  } else {
    fmt.Println("8 is odd")
  }

  if 8 % 4 == 0 {
    fmt.Println("8 is divisible by 4")
  }

  if 9 % 2 == 0 {
    fmt.Println("9 is enve")
  } else {
    fmt.Println("9 is odd")
  }

  if num := 9; num < 0 {
    fmt.Println(num, "is nagative")
  } else if num < 10 {
    fmt.Println(num, "has 1 digit")
  } else {
    fmt.Println(num, "has multiple digits")
  }
}
