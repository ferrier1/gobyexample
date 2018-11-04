// A panic typically means something went unexpectedly
// wrong. Mostly we use it to fail fast on errors that
// shouldn't occur during normal operation, or that we
// aren't prepared to handle gracefully.


package main

import "os"

func main() {

  panic("a problem")

  // common use of panic is to abort if a function
  // returns an error value that isnt known about or
  // how to handle it

  _, err := os.Create("/tmp/file")
  if err != nil {
    panic(err)
  }
}
