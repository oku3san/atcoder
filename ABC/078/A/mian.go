package main

import (
  "fmt"
  "strconv"
)

func main() {
  var X, Y string
  fmt.Scan(&X, &Y)

  hexX, _ := strconv.ParseInt(X, 16, 64)
  hexY, _ := strconv.ParseInt(Y, 16, 64)

  if hexX > hexY {
    fmt.Println(">")
  } else if hexX < hexY {
    fmt.Println("<")
  } else {
    fmt.Println("=")
  }
}
