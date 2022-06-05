package main

import "fmt"

// 各桁の和を計算する関数
func findSumOfDigits(n int) int {
  // 和の初期化
  sum := 0
  // 桁数が 0 になるまでループ
  for n > 0 {
    // n を 10 で割った余りを sum に足す
    sum += n % 10
    // n に n を 10 で割った値を入れる
    n /= 10
  }
  // 合計を返す
  return sum
}

func main() {
  // 標準入力から値受け取り
  // N は最大の整数、A 以上 B 以下の桁の和
  var N, A, B int
  fmt.Scan(&N, &A, &B)
  var total = 0
  // N の値までループ
  for i := 1; i <= N; i++ {
    // 各桁の合計を取得
    sum := findSumOfDigits(i)
    // 各桁の和の合計が A 以上 B 以下だったら合計に追加
    if sum >= A && sum <= B {
      total += i
    }
  }
  // 合計を表示
  println(total)
}
