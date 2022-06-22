package main

import (
  "fmt"
  "math/big"
)

func main() {
  var Y = big.NewInt(2)
  var N = new(big.Int)
  fmt.Scan(N)
  fmt.Println(N.Exp(Y, N, nil))
}
