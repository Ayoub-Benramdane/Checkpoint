package main

import "fmt"

func main() {
	Chunk([]int{}, 10)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 0)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 3)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 5)
	Chunk([]int{0, 1, 2, 3, 4, 5, 6, 7}, 4)
}

func Chunk(slice []int, size int) {
	parentSlice := [][]int{}
	newSlice := []int{}
	count := 0
	if size <= 0 {
		fmt.Println()
		return
	}
	for i := 0; i < len(slice); i++ {
		if i+size > len(slice) && count < len(slice) {
			newSlice = append(newSlice, count)
			count++
		} else if i >= count {
			for j := i; j < i+size; j++ {
				count++
				newSlice = append(newSlice, slice[j])
			}
			if newSlice != nil {
				parentSlice = append(parentSlice, newSlice)
			}
			newSlice = nil
		}
	}
	if newSlice != nil {
		parentSlice = append(parentSlice, newSlice)
	}
	fmt.Println(parentSlice)
}
