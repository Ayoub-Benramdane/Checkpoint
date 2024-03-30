package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 || args[0] == "" {
		fmt.Println()
		return
	}
	isFirst := true
	res := ""
	for _, c := range args[0] {
		if isFirst {
			if c != ' ' {
				res += string(c)
				isFirst = false
			}
		} else {
			if c == ' ' {
				res += string(c)
				isFirst = true
			} else {
				res += string(c)
			}
		}
	}
	fmt.Println(res)
}
