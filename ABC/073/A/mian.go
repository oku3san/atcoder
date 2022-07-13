package main

import (
  "fmt"
  "regexp"
  "strconv"
)

func main() {
  var N int
  fmt.Scan(&N)

  n := strconv.Itoa(N)

  match, _ := regexp.MatchString("9", n)
  if match {
    fmt.Println("Yes")
  } else {
    fmt.Println("No")
  }
}
