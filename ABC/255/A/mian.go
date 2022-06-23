package main

import "fmt"

func main() {
  var R, C int
  var a [2][2]int

  fmt.Scan(&R, &C)
  fmt.Scan(&a[0][0], &a[0][1])
  fmt.Scan(&a[1][0], &a[1][1])

  fmt.Println(a[R-1][C-1])
}
