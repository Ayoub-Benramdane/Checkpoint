package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		return
	}
	arg1 := args[0]
	arg2 := args[1]
	exist := true
	count := 0
	for i := 0; i < len(arg1); i++ {
		for j := count; j < len(arg2); j++ {
			count++
			if arg1[i] == arg2[j] {
				exist = true
				break
			} else {
				exist = false
			}
		}
	}
	if exist {
		z01.PrintRune('1')
		return
	}
	z01.PrintRune('0')
}
