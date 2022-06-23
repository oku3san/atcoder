package main

import "fmt"

func main() {
  var s string
  fmt.Scan(&s)

  // map[xxx:true yyy:false]
  a := make(map[rune]bool)

  // v には int が入る
  for _, v := range s {
    // すでに a[v] の配列が存在していたら狩猟
    if a[v] {
      fmt.Println("no")
      return
    }
    a[v] = true
  }
  fmt.Println("yes")
}
