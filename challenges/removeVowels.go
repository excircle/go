package main

import "fmt"
import "strings"

func removeVowels(s string) string {
	l := []string{"a", "e", "i", "o", "u"}
	for _, v := range l {
		s = strings.Replace(s, string(v), "", -1)
	}
	return s
}

func main() {
	testString := "This is a phrase"
	result := removeVowels(testString)
	fmt.Printf("%v\n%v\n", testString, result)
}
