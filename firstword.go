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
	res := ""
	for i := 0; i < len(args); i++ {
		if args[i] == ' ' {
			break
		}
		res += string(args[i])
	}
	fmt.Println(res)
}
