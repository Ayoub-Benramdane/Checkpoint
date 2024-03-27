package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	arg1 := os.Args[1]
	arg2 := os.Args[2]
	resFianl := make(map[string]bool)
	for _, char := range arg1 {
		for _, char1 := range arg2 {
			if char == char1 {
				if !resFianl[string(char)] {
					fmt.Printf(string(char))
					resFianl[string(char)] = true
				}
			}
		}
	}
}
