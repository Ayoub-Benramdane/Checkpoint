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
	count := 0
	count1 := 0
	for i := 0; i < len(arg); i++ {
		if arg[i] != ' ' && count == 0 {
			res += string(arg[i])
		} else if res != "" {
			count++
			if arg[i] != ' ' {
				count1 = 0
				resFinal += string(arg[i])
			} else {
				if count1 == 0 && resFinal != "" {
					count1++
					resFinal += " "
				}
			}
		}
	}
	resFinal += " " + res
	fmt.Println(resFinal)
}
