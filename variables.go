package main

import "fmt"

func main() {
  var a = "initial"
  fmt.Println(a)

  var b, c int = 1, 2
  fmt.Println(b, c)

  var d = true
  fmt.Println(d)

// non declared values are 0
  var e int
  fmt.Println(e)

  f := "short"
  fmt.Println(f)
}
