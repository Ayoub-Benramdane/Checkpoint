package main

import (
	"fmt"
	"os"
)

func main() {
	input := os.Args[1] + os.Args[2]
	res := make(map[string]bool)
	res1 := make(map[string]int)
	for _, char := range input {
		res1[string(char)]++
	}
	fmt.Println(res1)
	for _, char := range input {
		if !res[string(char)] {
			fmt.Printf(string(char))
			res[string(char)] = true
		}
	}
	fmt.Println()
}
