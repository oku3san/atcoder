package main

import "fmt"

func main() {
  var o, e string
  fmt.Scan(&o, &e)

  // o と e を順番に char で表示する
  for i := 0; i < len(e); i++ {
    fmt.Printf("%c%c", o[i], e[i])
  }
  // o の方が文字数が多い場合、o の文字数 -1 が配列の末尾の要素数になるので文字数-1の要素を表示
  if len(o) > len(e) {
    fmt.Printf("%c", o[len(o)-1])
  }
}
