package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	count := 0
	res := ""
	isFirst := true
	for i := 0; i < len(os.Args[1]); i++ {
		if isFirst && os.Args[1][i] == ' ' {
			continue
		}
		isFirst = false
		if os.Args[1][i] == ' ' && res != "" {
			count++
		} else if count != 0 {
			res += "   " + string(os.Args[1][i])
			count = 0
		} else {
			res += string(os.Args[1][i])
		}
	}
	fmt.Println(res)
}
