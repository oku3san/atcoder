package main

import (
  "fmt"
  "strings"
)

func main() {
  var s string
  fmt.Scan(&s)

  // 標準入力の値を逆転
  s = revers(s)

  // 逆転した文字列が、特定の文字列と合致するか
  // 合致した文字から削除していく
  // -1 を定義すると、合致する全てを置換する
  s = strings.Replace(s, "resare", "", -1)
  s = strings.Replace(s, "esare", "", -1)
  s = strings.Replace(s, "remaerd", "", -1)
  s = strings.Replace(s, "maerd", "", -1)

  var res string

  // すべて削除されたら
  if len(s) == 0 {
    res = "YES"
  } else {
    res = "NO"
  }
  fmt.Println(res)
}

// string の内容を逆転させる関数
// https://seven-901.hatenablog.com/entry/2021/06/14/234000
func revers(s string) string {
  runes := []rune(s)
  for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
    runes[i], runes[j] = runes[j], runes[i]
  }
  return string(runes)
}
