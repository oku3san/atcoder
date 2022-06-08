package main

import "fmt"

func main() {
  // n は枚数、y は合計金額
  var n, y int
  // 1行の標準入力から 2 つ値を取得
  fmt.Scan(&n, &y)

  // 10,000円を0枚からループ
  for i := 0; i <= n; i++ {
    // 5,000 円を0枚からループ、合計枚数から 10,000 円の枚数分引く
    for j := 0; j <= (n - i); j++ {
      // 合計枚数から 10,000 円と 5,000 円の枚数を引いて 1,000 円の枚数を取得
      k := n - (i + j)
      // 現時点での合計金額を取得
      total := 10000*i + 5000*j + 1000*k

      // 標準入力の値と合計が同じなら、print
      if total == y {
        fmt.Println(i, j, k)
        return
      }
    }
  }
  // 合計が合致しない場合の出力
  fmt.Println(-1, -1, -1)
}
