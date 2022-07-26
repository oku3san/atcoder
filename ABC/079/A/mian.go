package main

import (
  "fmt"
  "strconv"
)

func main() {
  var X int
  fmt.Scan(&X)

  x := strconv.Itoa(X)

  if x[0] == x[1] && x[1] == x[2] || x[1] == x[2] && x[2] == x[3] {
    fmt.Println("Yes")
  } else {
    fmt.Println("No")
  }
}
