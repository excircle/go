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

# FMT's "Stringer" Interface

The FMT package implements an interface called "[Stringer](https://pkg.go.dev/fmt#Stringer)" which can be view at the link provided.

This interface allows Golang developers to create a formalized "Print" method for any new types that may need to be printed.

### Without "String()" method

```go
package main

import "fmt"

type order struct {
	ProductCode int
	Quantity float64
	Status orderStatus
}

var orders = []order{}

var rawOrders = []string{
	`{"productCode": 1111, "quantity": 5, "status": 1}`,
	`{"productCode": 2222, "quantity": 42.3, "status": 1}`,
	`{"productCode": 3333, "quantity": 19, "status": 1}`,
	`{"productCode": 4444, "quantity": 8, "status": 1}`,
}

for _, rawOrder := range rawOrders {
	var newOrder order
	err := json.Unmarshal([]byte(rawOrder), &newOrder)
	if err != nil {
		log.Print(err)
		continue
	}
	orders = append(orders, newOrder)
}

fmt.Println(orders)
/*
OUTPUTS:
[{1111 5 1} {2222 42.3 1} {3333 19 1} {4444 8 1}]
```

### WITH "String()" method

```go
package main

import "fmt"

type order struct {
	ProductCode int
	Quantity float64
	Status orderStatus
}
// By adding the below string method to the "order" type, we are able
// to have a formated print string that is automatically utilized by
// the fmt package
func (o order) String() string {
	return fmt.Sprintf("Product code; %v, Quantity: %v, Status: %v\n",
		o.ProductCode, o.Quantity, orderStatusToText(o.Status))
}

var orders = []order{}

var rawOrders = []string{
	`{"productCode": 1111, "quantity": 5, "status": 1}`,
	`{"productCode": 2222, "quantity": 42.3, "status": 1}`,
	`{"productCode": 3333, "quantity": 19, "status": 1}`,
	`{"productCode": 4444, "quantity": 8, "status": 1}`,
}

for _, rawOrder := range rawOrders {
	var newOrder order
	err := json.Unmarshal([]byte(rawOrder), &newOrder)
	if err != nil {
		log.Print(err)
		continue
	}
	orders = append(orders, newOrder)
}

fmt.Println(orders)
/*
OUTPUTS:
[Product code; 1111, Quantity: 5, Status: new
 Product code; 2222, Quantity: 42.3, Status: new
 Product code; 3333, Quantity: 19, Status: new
 Product code; 4444, Quantity: 8, Status: new
]
```
