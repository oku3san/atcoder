package main

import (
  "fmt"
  "strings"
)

func main() {
  var s string
  fmt.Scan(&s)
  fmt.Printf("%d", strings.Count(s, "1"))
}
