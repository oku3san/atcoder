package main

import (
  "fmt"
)

func main() {
  var X, A, B int
  fmt.Scan(&X, &A, &B)

  if B-A <= 0 {
    fmt.Println("delicious")
  } else if (B - A) <= X {
    fmt.Println("safe")
  } else {
    fmt.Println("dangerous")
  }
}
