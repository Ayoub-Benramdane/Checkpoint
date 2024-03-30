package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	count := 0
	res := ""
	isFirst := true
	for i := 0; i < len(args[0]); i++ {
		if args[0][i] != ' ' {
			isFirst = false
		}
		if args[0][i] == ' ' && res != "" {
			count++
			continue
		} else if count != 0 && isFirst == false {
			res += "   " + string(args[0][i])
			count = 0
		} else if count == 0 && args[0][i] != ' ' {
			res += string(args[0][i])
		}
	}
	fmt.Println(res)
}
