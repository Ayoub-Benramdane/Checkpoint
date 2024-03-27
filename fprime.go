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
	num, err := strconv.Atoi(args[0])
	if err != nil || num <= 1 {
		return
	}
	primes(num)
}

func primes(n int) {
	for n%2 == 0 {
		if n/2 != 1 {
			fmt.Print("2*")
		} else {
			fmt.Print("2")
		}
		n /= 2
	}
	for i := 3; i <= n; i += 2 {
		if n%i == 0 {
			if n/i != 1 {
				fmt.Printf("%d*", i)
				n /= i
				i = 3
			}
		}
	}
	if isPrime(n) {
		fmt.Println(n)
		return
	}
}

func isPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	if nb == 2 {
		return true
	}
	if nb%2 == 0 {
		return false
	}
	for i := 3; i <= nb/3; i += 2 {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
