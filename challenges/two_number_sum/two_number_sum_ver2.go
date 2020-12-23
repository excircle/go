package main

import (
	"fmt"
	"sort"
)

func TwoNumberSum(array []int, target int) []int {
	sort.Ints(array)
	left := 0
	right := len(array) - 1
	for left < right {
		potential := array[left] + array[right]
		if potential == target {
			return []int{array[left], array[right]}
		} else if potential < target {
			left += 1
		} else {
			right -= 1
		}
	}
	return []int{}
}

//Find two numbers that add up to 9
func main() {
	a := []int{3, 5, -4, 8, 11, 1, -1, 6}
	t := 9
	fmt.Println(TwoNumberSum(a, t))
}
