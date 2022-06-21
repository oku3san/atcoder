package main

import "fmt"

// 配列野中に指定の文字があるか確認
func existsCheck(slice []int, target int) bool {
  for _, num := range slice {
    if num == target {
      return true
    }
  }
  return false
}

func main() {
  var x, y int
  fmt.Scan(&x, &y)

  var a = []int{1, 3, 5, 7, 8, 10, 12}
  var b = []int{4, 6, 9, 11}
  var c = []int{2}

  if existsCheck(a, x) && existsCheck(a, y) || existsCheck(b, x) && existsCheck(b, y) || existsCheck(c, x) && existsCheck(c, y) {
    fmt.Println("Yes")
  } else {
    fmt.Println("No")
  }

}
