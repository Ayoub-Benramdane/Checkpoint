package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args[1:]) != 1 {
		return
	}
	arg := os.Args[1]
	res := ""
	resFinal := ""
	for i := 0; i < len(arg); i++ {
		if  i == len(arg)-1 {
			res += string(arg[i])
			resFinal = res + " " + resFinal
		} else if arg[i] != ' ' {
			res += string(arg[i])
		} else if res != "" {
			resFinal = res + " " + resFinal
			res = ""
		}
	}
	fmt.Println(resFinal)
}
