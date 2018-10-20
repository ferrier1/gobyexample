package main

import "fmt"

func vals() (int, int, int) {
  return 3, 7, 5
}

func main() {

  a, b, d := vals()
  fmt.Println(a)
  fmt.Println(b)
  fmt.Println(d)

  _, c, _ :=  vals()
  fmt.Println(c)

}
