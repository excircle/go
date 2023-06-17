# The FMT Library

# General Information

The FMT library has many functions. This section convers some of the material that is contained within the library.

### FMT returns two values

```go
package main

import "fmt"

func main() {
	x := map[int]string{
		1: "alex",
		2: "nico"}

	_, ok := x[2]

	a, b := fmt.Println(ok)
	fmt.Println(a, b)
    // Returns number of bytes writen, and any errors encoutered
}
```
