package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	base := "0123456789abcdef"
	args := os.Args[1:]
	if len(args) != 1 {
		return
	}
	res := ""
	arg, _ := strconv.Atoi(args[0])
	for _, c := range args[0] {
		if c < '0' || c > '9' {
			fmt.Println("ERROR")
			return
		}
	}
	for arg != 0 {
		if arg%len(base) != 0 {
			res = string(base[arg%len(base)]) + res
			arg /= len(base)
		} else {
			z01.PrintRune(rune(base[arg/len(base)]))
			z01.PrintRune('0')
			break
		}
	}
	fmt.Println(res)
}
