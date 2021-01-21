package main

import "fmt"

func IsValidSubsequence(array []int, sequence []int) bool {
	if len(array) == len(sequence) {
		for i := 0; i < len(array); i++ {
			if array[i] != sequence[i] {
				return false
			}
			return true
		}
	}
	intsPresent := map[int]bool{}
	seqExists := true
	for _, v := range array {
		intsPresent[v] = true
	}
	for _, v := range sequence {
		if intsPresent[v] != true {
			seqExists = false
		}
	}
	return seqExists
}

func main() {
	a := []int{5, 1, 22, 25, 6, -1, 8, 10}
	s := []int{5, 1, 22, 23, 6, -1, 8, 10}
	result := IsValidSubsequence(a, s)
	fmt.Println(result)
}
