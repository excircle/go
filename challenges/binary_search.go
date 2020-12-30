package main

import "fmt"

func BinarySearch(array []int, target int) int {
	left 		:= 0
	right		:= len(array)-1
	iterations	:= len(array)-1
	var median int
	for i := 0; i < iterations; i++ {
		median = (right + left) / 2
		if array[median] == target {
			return median
		}
		if array[median] > target {
			right = median - 1
		}
		if array[median] < target {
			left = median + 1
		}
	}
	return -1
}

func main()  {
	myArray	:= []int{1,5,23,111}
	target 	:= 23
	results := BinarySearch(myArray, target)
	fmt.Println(results)
}