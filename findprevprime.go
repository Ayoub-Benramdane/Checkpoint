package main

import "fmt"

func main() {
	fmt.Println(FindPrevPrime(225225))
	fmt.Println(FindPrevPrime(28))
}

func FindPrevPrime(nb int) int {
	if !IsPrime(nb) {
		nb = FindPrevPrime(nb - 1)
	}
	return nb
}

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	for i := 2; i*i <= nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
