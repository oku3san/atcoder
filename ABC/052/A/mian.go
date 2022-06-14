package main

import (
  "fmt"
)

func main() {
  var A, B, C, D int
  fmt.Scan(&A, &B, &C, &D)

  first := A * B
  second := C * D
  ans := 0

  switch {
  case first > second:
    ans = first
  case first < second:
    ans = second
  case first == second:
    ans = first
  }
  fmt.Println(ans)
}
