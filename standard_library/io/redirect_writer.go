package main

import (
  "bytes"
  "io"
  "os"
)

//###############
//###io.Writer###
//###############
//
//io.Writer is an interface for the Write() method.
//In this example, we create a function "myFunction" which works with an io.Writer interface
//The main demonstrates the ability to pass "myFunction"
//  both os.Stdout and bytes.Buffer type, and write to them using our array of bytes.
//os.Stdout returns a File, which effectively points to "/dev/stdout," and has a Write() method.
//byes.Buffer also has a Write() method.
//Through the io.Writer interface, myFunction can call the Write() method on both bytes.Buffer & os.Stdout types.

func myFunction(outputDirection io.Writer) bool {
  outputDirection.Write([]byte("Terminal message for my users\n"))
  return true
}

func main() {
  //In production, perhaps you'll want all output to stdout
  myFunction(os.Stdout)

  //In *_test.go files, to ensure "myFunction" returns true without spamming the terminal
  var b bytes.Buffer
  myFunction(&b)

  //and if I need to look at or inspect something in test,
  //do something like this
  b.WriteTo(os.Stdout)
}
