package main

import "fmt"

func main() {
  var A, B, C, X int
  fmt.Scan(&A, &B, &C, &X)
  res := 0

  for a := 0; a <= A; a++ {
    for b := 0; b <= B; b++ {
      for c := 0; c <= C; c++ {
        total := 500*a + 100*b + 50*c
        if total == X {
          res++
        }
      }
    }
  }
  fmt.Println(res)
}
