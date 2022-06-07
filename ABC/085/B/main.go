package main

import "fmt"

func main() {
  // 枚数を取得
  var n int
  fmt.Scan(&n)
  // 枚数分だけ配列作成
  d := make([]int, n)
  // 標準入力から枚数ごとの直径を取得
  for i := 0; i < n; i++ {
    fmt.Scan(&d[i])
  }

  // 連想配列を作成
  // map[6:1 8:2 10:1] のような形になる
  p := make(map[int]int)
  // 枚数の分ループ
  for v := range d {
    // d 配列の value(直径) ごとに key(個数) の添字を増やしている
    // d の値が同じ場合、添字が増える
    p[d[v]]++
  }
  // map の要素数を取得
  fmt.Println(len(p))
}
