package main

import (
  "fmt"
  "sort"
)

func main() {
  // 標準入力2行分
  var n, a int
  // 1行目を取得
  fmt.Scan(&n)
  // 1行目の数分、arr 配列を広げる
  arr := make([]int, n)

  // 配列の数だけ、2行目の値を配列に格納する
  for i := range arr {
    fmt.Scan(&a)
    arr[i] = a
  }

  // 降順ソート
  sort.Sort(sort.Reverse(sort.IntSlice(arr)))

  // alice と bob のそれぞれの得点格納用
  var alice, bob int

  // 配列の数(カードの数)分ループ
  for i := range arr {
    // 奇数回は alice の得点、偶数回は bob の得点
    if i%2 == 0 {
      alice += arr[i]
    } else {
      bob += arr[i]
    }
  }
  // alice と bob の得点差を表示
  fmt.Println(alice - bob)
}
