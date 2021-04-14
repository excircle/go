package main

import "fmt"

func ReverseInts(s []int) {
  first := 0
  last := len(s)-1
  for first < last {
    s[first], s[last] = s[last], s[first]
    first++
    last--
  }
}

func main() {
  x := []int{1,2,3,4}
  fmt.Println(x)
  ReverseInts(x)
  fmt.Println(x)
}