package main

import "fmt"

func TwoNumberSum(array []int, target int) []int {
	for k1 := 0; k1 < len(array)-1; k1++ {
		for k2 := k1 + 1; k2 < len(array); k2++ {
			v1 := array[k1]
			v2 := array[k2]
			fmt.Printf("{%d, %d}:[%d, %d]\n", k1, k2, v1, v2)
			if v1+v2 == target {
				return []int{v1, v2}
			}
		}
	}
	return []int{}
}

func main() {
	a := []int{3, 5, -4, 8, 11, 1, -1, 6}
	t := 1000
	fmt.Println(TwoNumberSum(a, t))
}
