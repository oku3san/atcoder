package main

import "fmt"

func main() {
  var a, b, result string
  fmt.Scan(&a, &b)

  if a == "H" {
    if b == "H" {
      result = "H"
    } else {
      result = "D"
    }
  } else {
    if b == "H" {
      result = "D"
    } else {
      result = "H"
    }
  }

  fmt.Println(result)
}
