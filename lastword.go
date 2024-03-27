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
	resFinal := ""
	for i := 0; i < len(args); i++ {
		if args[i] == ' ' {
			res = ""
			continue
		}
		res += string(args[i])
		resFinal = res
	}
	fmt.Println(resFinal)
}
