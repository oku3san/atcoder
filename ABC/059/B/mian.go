package main

import (
  "fmt"
  "math/big"
)

func main() {
  //var A, B float64
  //fmt.Scan(&A, &B)
  //switch {
  //case A == B:
  //  fmt.Println("EQUAL")
  //case A > B:
  //  fmt.Println("GREATER")
  //case A < B:
  //  fmt.Println("LESS")
  //}

  var A = new(big.Int)
  var B = new(big.Int)
  fmt.Scan(A, B)
  switch A.Cmp(B) {
  case 0:
    fmt.Println("EQUAL")
  case 1:
    fmt.Println("GREATER")
  case -1:
    fmt.Println("LESS")
  }

}
