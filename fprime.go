package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		return
	}
	n, err := strconv.Atoi(args[0])
	if err != nil || n <= 1 {
		return
	}
	Fprime(n)
}

func Fprime(n int) {
	for i := 2;n/i != 0;i++ {
		if n%i == 0 {
			if n/i != 1 {
				fmt.Printf("%d*", i)
			} else {
				fmt.Printf("%d", i)
			}
			n /= i
			i--
		}
	}
	fmt.Println()
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
