package main

import (
  "fmt"
  "sort"
)

func main() {
  // 標準入力
  var n, a int
  fmt.Scan(&n)
  // arr 配列を広げる
  arr := make([]int, n)

  // 標準入力の値を配列に格納する
  for i := range arr {
    fmt.Scan(&a)
    arr[i] = a
  }

  // 昇順ソート
  sort.Ints(arr)
  fmt.Println(arr[n-1] - arr[0])
}
