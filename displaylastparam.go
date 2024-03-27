package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	for i := 0; i < len(args); i++ {
		fmt.Println(args[len(args)-1])
		break
	}
}
