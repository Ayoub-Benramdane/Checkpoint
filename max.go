package main

import (
	"fmt"
)

func main() {
	a := []int{23, 123, 1, 11, 55, 93}
	max := Max(a)
	fmt.Println(max)
}

func Max(a []int) int {
	max := a[0]
	for i := 0; i < len(a)-1; i++ {
		for j := i+1; j < len(a); j++ {
			if max < a[j] {
				max = a[j]
			}
		}
	}
	return max
}
