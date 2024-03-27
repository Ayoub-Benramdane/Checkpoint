package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	args := os.Args[1]
	for i := 0; i < len(args); i++ {
		if args[i] >= 'a' && args[i] <= 'z' {
			for j := 0; j <= int(args[i]-97); j++ {
				fmt.Printf(string(args[i]))
			}
		} else if args[i] >= 'A' && args[i] <= 'Z' {
			for j := 0; j <= int(args[i]-65); j++ {
				fmt.Printf(string(args[i]))
			}
		} else {
			fmt.Printf(string(args[i]))
		}
	}
}
