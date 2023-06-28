# Go Structs

# How to Utilize Go Fields In Code

The following is an example of how to use Go structs in fields.

```go
package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Fname  string
	Lname  string
	City   string
	Mobile int64
}

func main() {
	s := Student{"Chetan", "Tulsyan", "Bangalore", 7777777777}
	v := reflect.ValueOf(s)
	typeOfS := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("%s:\t%v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
	}
}

```
