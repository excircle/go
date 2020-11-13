package main

import "time"
import "fmt"

// 1,000,000,000 seconds
const gigasecond = 1000000000

func main() {
  // time.Now() returns a time.Time struct. Looks like this when Println'ed:
  // "2009-11-10 23:00:00 +0000 UTC m=+0.000000001"
  // The values for things like year and day are populated and assembled into the struct from the Go "runtime" package
  here := time.Now()

  // The "Add" method of the time.Time struct accepts an int64 type called "Duration"
  // The method definition is as below
  // func (t Time) Add(d Duration) Time {}
  // Add() modifies the attributes of struct "t Time" by returning t Time + Duration (which is an int64)
  // Here, we multiply our const "gigasecond" by the const inside package time called "time.Second"
  // The "time.Second" const is: "Second = 1000 * Millisecond"
  future := here.Add(gigasecond*time.Second)
  past := here.Add(-gigasecond*time.Second)
  fmt.Println("PAST:", past)
  fmt.Println("PRESENT:", here)
  fmt.Println("FUTURE:", future)
}

