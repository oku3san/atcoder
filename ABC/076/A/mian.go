package main

import "fmt"

func main() {
  var R, G int
  fmt.Scan(&R, &G)
  fmt.Println(2*G - R)
}

/*
(R + X) / 2 = G
R + X = 2G
X = 2G - R

より、2G-Rを答える。
*/
