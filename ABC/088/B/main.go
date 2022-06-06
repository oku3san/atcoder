package main

import (
  "fmt"
  "sort"
)

func main() {
  var n, a int
  fmt.Scan(&n)
  arr := make([]int, n)
  for i := range arr {
    fmt.Scan(&a)
    arr[i] = a
  }
  sort.Sort(sort.Reverse(sort.IntSlice(arr)))

  var alice, bob int
  for i := range arr {
    if i%2 == 0 {
      alice += arr[i]
    } else {
      bob += arr[i]
    }
  }
  fmt.Println(alice - bob)
}
