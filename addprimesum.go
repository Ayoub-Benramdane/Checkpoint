package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	prime := true
	res := []int{}
	resFinal := 0
	if len(args) != 1 {
		fmt.Println(resFinal)
		return
	}
	arg, _ := strconv.Atoi(args[0])
	if arg < 0 {
		fmt.Println(resFinal)
		return
	}
	for i := arg; i >= 2; i-- {
		for j := 2; j < i; j++ {
			if i%j == 0 {
				prime = false
				break
			}
		}
		if prime {
			res = append(res, i)
		}
		prime = true
	}
	for i := 0; i < len(res); i++ {
		resFinal += res[i]
	}
	fmt.Println(res)
	fmt.Println(resFinal)
}
