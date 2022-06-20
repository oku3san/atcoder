package main

import (
  "fmt"
  "strings"
)

func main() {
  var s1, s2, s3 string
  var concat string
  fmt.Scan(&s1, &s2, &s3)
  concat = s1[0:1] + s2[0:1] + s3[0:1]
  fmt.Println(strings.ToUpper(concat))
}
