package main

import (
  "fmt"
  "strconv"
)

func main() {
  //var r, g, b int
  //fmt.Scan(&r, &g, &b)
  //
  //score := (r * 100) + (g * 10) + (b * 1)
  //
  //if score%4 == 0 {
  //  fmt.Println("YES")
  //} else {
  //  fmt.Println("NO")
  //}

  var r, g, b string
  fmt.Scan(&r, &g, &b)

  i, _ := strconv.Atoi(r + g + b)

  if i%4 == 0 {
    fmt.Println("YES")
  } else {
    fmt.Println("NO")
  }
}
