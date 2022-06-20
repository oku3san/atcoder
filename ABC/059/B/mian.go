package main

import "fmt"

func main() {
  var A, B float64
  fmt.Scan(&A, &B)
  switch {
  case A == B:
    fmt.Println("EQUAL")
  case A > B:
    fmt.Println("GREATER")
  case A < B:
    fmt.Println("LESS")
  }
}
