package main

import "fmt"

func main() {
  var A, B, C int

  fmt.Scan(&A, &B, &C)

  switch {
  case A == B:
    fmt.Println(C)
  case A == C:
    fmt.Println(B)
  case B == C:
    fmt.Println(A)
  }
}
