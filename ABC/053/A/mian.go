package main

import "fmt"

func main() {
  var x int
  fmt.Scan(&x)

  var result string
  if x >= 1200 {
    result = "ARC"
  } else {
    result = "ABC"
  }
  fmt.Println(result)
}
