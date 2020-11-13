package main

import "time"
import "fmt"

// I am not a fan of time parsing in Go.
// As I understand it, you must define a layout.
// I define the layout here as a const
const layoutISO = "2006-01-02"

func main() {
	// variable "date" represents the string I want to parse
	// If I wanted to turn this string into a time.Time object for evaluation
	// I can use the Parse() method to return a time.Time object and an Err object
	// func Parse(layout, value string) (Time, error)
	// The docs do not explain what layout has to be
	// However, I just use a string type as a const and it works
	date := "1999-12-31"
	t, _ := time.Parse(layoutISO, date)
	// The following Println returns a time.Time object
	fmt.Println(t)
	// Println="1999-12-31 00:00:00 +0000 UTC"
}
