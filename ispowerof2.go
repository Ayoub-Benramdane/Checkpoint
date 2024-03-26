package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	res, _ := strconv.Atoi(args[1])
	if t(res) {
		fmt.Println("True")
	} else {
		fmt.Println("False")

	}
}

func t(n int) bool {

	for n != 2 {
		if n%2 != 0 {
			return false
		}
		n /= 2
	}
	return true
}
