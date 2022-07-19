package main

import "fmt"

func main() {
  var C1, C2 string
  fmt.Scan(&C1, &C2)

  if C1[0] == C2[2] && C1[1] == C2[1] && C1[2] == C2[0] {
    fmt.Println("YES")
  } else {
    fmt.Println("NO")
  }
}
