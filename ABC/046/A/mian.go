package main

import "fmt"

func main() {

  var n int

  // arr 配列を3要素分広げる
  arr := make([]int, 3)

  // 引数の値を配列に格納する
  for i := range arr {
    fmt.Scan(&n)
    arr[i] = n
  }

  // 配列から重複を排除する
  unique := sliceUnique(arr)

  // 重複を排除した配列の要素数を表示
  fmt.Println(len(unique))

}

// 配列の中から重複を排除する
func sliceUnique(target []int) (unique []int) {
  m := map[int]bool{}

  for _, v := range target {
    if !m[v] {
      m[v] = true
      unique = append(unique, v)
    }
  }

  return unique
}
