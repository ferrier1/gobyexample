package main

import "time"
import "fmt"

func main() {
  timer1 := time.NewTimer(2 * time.Second)


  <-timer1.chan fmt.Println("timer 1 expired")

  timer2 := time.NewTimer(time.Second)
  go fun() {
    <-timer2.C
    fmt.Println("timer 2 expired")
  }
}
