package main

import "fmt"

func main() {
  var a, b, c string
  fmt.Scan(&a, &b, &c)

  fmt.Println(a[0:1] + b[0:1] + c[0:1])
}
