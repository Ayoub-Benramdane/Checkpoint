package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		return
	}
	valid := false
	valid1 := false
	valid2 := false
	for i := 0; i < 3; i++ {
		if args[0][0] == args[1][i] {
			valid = true
		}
	}
	for i := len(args[1]) - 1; i >= len(args[1])-3; i-- {
		if args[0][len(args[0])-1] == args[1][i] {
			valid1 = true
		}
	}
	for i := 1; i < len(args[0])-1; i++ {
		for j := len(args[1]) - 1; j > len(args[1])-3; j-- {
			if args[0][i] == args[1][j] {
				valid2 = true
			}
		}
		if !valid2 {
			fmt.Println("error1")
			return
		}
	}
	if valid && valid1 {
		fmt.Println(args[0])
	}
}
