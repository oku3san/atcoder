package main

import "fmt"

func main() {
  var alice, bob int
  fmt.Scan(&alice, &bob)
  var result string

  switch {
  case alice == bob:
    result = "Draw"
  case alice == 1:
    result = "Alice"
  case bob == 1:
    result = "Bob"
  case alice > bob:
    result = "Alice"
  case alice < bob:
    result = "Bob"
  }
  fmt.Println(result)
}
