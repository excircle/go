# Reflect Package

This directory is dedicated to the Go standard library "Reflect."

# Reflecting Type Data

```go
package main

import (
	"context"
	"fmt"
	"log"
	"reflect"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

func exposeObject(x aws.Config) {
	v := reflect.ValueOf(x)
	typeOfS := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("%s:\t %v\n\n", typeOfS.Field(i).Name, v.Field(i).Interface())
	}
}

func main() {
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	exposeObject(cfg)
}
```
