package main

import (
	"fmt"
)

func primes(n int) {
	for n%2 == 0 {
		fmt.Print("2 * ")
		n /= 2
	}

	for i := 3; i <= n; i += 2 {
		if n%i == 0 {
			fmt.Printf("%d * ", i)
			//	fmt.Print("n: ", n, " " )
			n /= i
			i = 3

		}
	}
	if isPrime(n) && n > 2 {
		fmt.Println(n)
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
