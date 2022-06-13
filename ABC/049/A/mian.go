package main

import (
  "fmt"
  "regexp"
)

func main() {
  var c string
  fmt.Scan(&c)

  var result string
  match, _ := regexp.MatchString("[a|i|u|e|o]", c)
  if match {
    result = "vowel"
  } else {
    result = "consonant"
  }
  fmt.Println(result)

  //if strings.Contains("aiueo", c) {
  //  fmt.Println("vowel")
  //} else {
  //  fmt.Println("consonant")
  //}
}
