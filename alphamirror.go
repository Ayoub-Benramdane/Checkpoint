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
			fmt.Printf(string('a' + ('z' - args[i])))
		} else if args[i] >= 'A' && args[i] <= 'Z' {
			fmt.Printf(string('A' + ('Z' - args[i])))
		} else {
			fmt.Printf(string(args[i]))
		}
	}
}
