package main

import (
	"fmt"
	"sort"
)

func ThreeNumberSum(array []int, target int) [][]int {
	sort.Ints(array)
	res := [][]int{}
	var answer []int
	var curSum int

	for i := 0; i != len(array)-2; i++ {
		l := i+1
		r := len(array)-1
		for l < r {
			curSum = array[i] + array[l] + array[r]
			if curSum == target {
				answer = []int{array[i], array[l], array[r]}
				res = append(res, answer)

				l++
				r--
			}
			if curSum < target {
				l++
			}
			if curSum > target {
				r--
			}
		}
	}
	return res
}

func main() {
	input := []int{12,3,1,2,-6,5,-8,6}
	target := 0
	results := ThreeNumberSum(input, target)
	fmt.Println(results)
}