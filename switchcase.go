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
			fmt.Printf(string(args[i]-32))
		} else if args[i] >= 'A' && args[i] <= 'Z' {
			fmt.Printf(string(args[i]+32))
		} else {
			fmt.Printf(string(args[i]))
		}
	}
}
