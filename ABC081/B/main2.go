// ABC 081 B - Shift Only
package main

import "fmt"

func main() {
  var n, cnt int
  // 1回目の入力の数文の配列を作る用
  var a []int

  // 整数の数を取得
  fmt.Scan(&n)
  // 整数の数分配列を広げる(slice)
  a = make([]int, n)

  // 整数の数分 for ループをまわし、空白で区切られた数をa配列に入れていく
  for i := 0; i < n; i++ {
    fmt.Scan(&a[i])
  }

  flag := true
  // flag が false になるまで for ループ
  for flag {
    // 配列の数分 for ループを回す
    for i := 0; i < n; i++ {
      // a[i]の整数が偶数なら2で割ったものに置き換え
      if a[i]%2 == 0 {
        a[i] /= 2
      } else { // 偶数じゃないものがあれば終わり
        flag = false
      }
    }
    // a[] のすべてが偶数だった数をカウント
    if flag {
      cnt++
    }
  }
  fmt.Print(cnt)
}
