package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	res := ""
	if len(args) != 4 {
		return
	}
	for i := 0; i < len(args[1]); i++ {
		if string(args[1][i]) == args[2] {
			res += string(args[3])
		} else {
			res += string(args[1][i])
		}
	}
	fmt.Println(res)
}
