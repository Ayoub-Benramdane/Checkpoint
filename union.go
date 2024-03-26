package main

import (
	"fmt"
	"os"
)

func main() {
	input := os.Args[1] + os.Args[2]
	arr := make(map[string]int)
	res := make(map[string]bool)

	for _, char := range input {
		arr[string(char)]++
	}
	fmt.Println(arr)

	for _, char := range input {
		if !res[string(char)] {
			fmt.Printf(string(char))
			res[string(char)] = true
		}
	}
	fmt.Println()
}
