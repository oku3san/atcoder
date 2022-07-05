package main

import "fmt"

func main() {
  var X, t int
  fmt.Scan(&X, &t)

  if X-t > 0 {
    fmt.Print(X - t)
  } else {
    fmt.Println(0)
  }
}
