package main

import "fmt"

func main() {
  var n int
  fmt.Scan(&n)

  paid := n * 800
  back := 200 * (n / 15)

  fmt.Println(paid - back)

}
