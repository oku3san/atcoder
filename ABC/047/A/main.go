package main

import (
  "fmt"
  "sort"
)

func main() {
  // 標準入力
  var n int
  // arr 配列を広げる
  arr := make([]int, 3)

  // 標準入力の値を配列に格納する
  for i := range arr {
    fmt.Scan(&n)
    arr[i] = n
  }

  // 降順ソート
  sort.Sort(sort.IntSlice(arr))

  var result string
  if arr[0]+arr[1] == arr[2] {
    result = "Yes"
  } else {
    result = "No"
  }
  fmt.Println(result)
}
