# Lessons on Pointers

Pointers are often a "point" of confusion when discussing languages like C or Go.

This document seeks to clarify the rationale behind pointers.

# Basic Usage of a Pointer

Every function, whether it is the `main()` function or not, has a variable scope that is inaccessible to other functions without pointers.

In order for us to be able to access the memory addresses of variables from another functions scope, we must use pointers.

```go
package main

import "fmt"

func increment(a *int) { // Pointer to an int
	*a++ // Follow or "dereference" the variable 'a' and increment (++) the value it relates to
}

func main() {
	x := 5
	fmt.Println(x)

	increment(&x) // Memory address of 'X' from 'main()'
	fmt.Println(x)
}
```

# Pointer Considerations

The only time you'll want to use a pointer is if you're dealing with a function that expects to handle a interface.

For example, when dealing with a JSON document in Go, the methods which are exposed by the JSON interface must be given directly to to the function.

```go
f := struct {
Name string `json:"name"`
Age int `json:"age"`
}
err := json.Unmarshal([]byte(`{"name": "Bob", "age": 30}`), &f)
```

In the above example, we would potentially be dealing with the wrong data if we were calling JSON methods on a copy of `f` versus the exact memory address of `f`.

# Pointers When Using Maps

In Go, Slices and Maps use "Array" types behind the scenes. The associated arrays are references in the Slice and Map type via pointers.

<b>IMPORTANT:</b> Because Maps and Slices use pointers to access their internal arrays, Maps and Slices impact their stored values directly when used in functions. This is due to the pointers which are unavoidably references during the passing of themselves into functions.
