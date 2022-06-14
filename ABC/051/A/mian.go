package main

import (
  "fmt"
  "strings"
)

func main() {
  var c string
  fmt.Scan(&c)

  replaced := strings.Replace(c, ",", " ", -1)

  fmt.Println(replaced)
}
