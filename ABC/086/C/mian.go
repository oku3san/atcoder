package main

import (
  "fmt"
  "math"
)

func main() {
  // 点の数を取得
  var n int
  fmt.Scan(&n)

  // 時刻、(x,y)を取得
  var t int
  var x, y float64 // math.Abs は float を使う
  t, x, y = 0, 0.0, 0.0

  // 次の目的地用
  var tNext int
  var xNext, yNext float64 // math.Abs は float を使う

  flag := "Yes"
  // 点の数分ループ
  for i := 0; i < n; i++ {
    // 次の点の値を取得
    fmt.Scan(&tNext, &xNext, &yNext)

    // 残り時間
    T := tNext - t

    // 値が0からどれだけ離れているかの距離を表す絶対値を計算
    dist := math.Abs(xNext-x) + math.Abs(yNext-y)

    // 残り時間より必要な移動距離が長い場合
    if T < int(dist) {
      flag = "No"
    }

    // 到着後の余り時間が偶数なら行き帰りできる
    if (T-int(dist))%2 == 1 {
      flag = "No"
    }

    // 値を更新する
    t, x, y = tNext, xNext, yNext
  }
  fmt.Println(flag)
}
