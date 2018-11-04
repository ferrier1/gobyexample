// Go's `sort` package implements sorting for builtins
// and user-defined types. We'll look at sorting for
// builtins first.

package main

import (
  "fmt"
  "sort"
)

func main() {

  // sorting strings example
  // sorting is in-place so it changes the given value
  // and does not return new

  strs := []string{"c", "a", "b"}
  sort.Strings(strs)
  fmt.Println("Strings:", strs)

  // example of sorting ints

  ints := []int{7, 2, 5}
  sort.Ints(ints)
  fmt.Println("Ints:   ", ints)


  // sort can also be used to check if a slice is
  // already in sorted order

  s := sort.IntsAreSorted(ints)
  fmt.Println("Sorted: ", s)
}
