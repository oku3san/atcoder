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

  //    var N float64
  //    fmt.Scan(&N)
  //    fmt.Println(int(math.Pow(2,N)))
}
