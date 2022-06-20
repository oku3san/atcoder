package main

import (
  "fmt"
  "strings"
)

func main() {
  var A, B, C, result string
  fmt.Scan(&A, &B, &C)

  if strings.HasSuffix(A, B[0:1]) && strings.HasSuffix(B, C[0:1]) {
    result = "YES"
  } else {
    result = "NO"
  }
  fmt.Println(result)
}
