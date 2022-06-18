package main

import "fmt"

func main() {
  var a, b, c int
  fmt.Scan(&a, &b, &c)

  var result string
  if b-a == c-b {
    result = "YES"
  } else {
    result = "NO"
  }

  fmt.Println(result)
}
