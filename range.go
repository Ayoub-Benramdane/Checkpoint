package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		return
	}
	arg1, err1 := strconv.Atoi(args[0])
	arg2, err2 := strconv.Atoi(args[1])
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	if arg1 >= arg2 {
		for i := arg1; i >= arg2; i-- {
			if i > arg2 {
				fmt.Printf(strconv.Itoa(i) + " ")
			} else {
				fmt.Printf(strconv.Itoa(i))
			}
		}
	} else {
		for i := arg1; i <= arg2; i++ {
			if i < arg2 {
				fmt.Printf(strconv.Itoa(i) + " ")
			} else {
				fmt.Printf(strconv.Itoa(i))
			}
		}
	}
}
