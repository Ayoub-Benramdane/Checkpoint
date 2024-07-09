package main

import "fmt"

func main() {
	fmt.Println(FindPrevPrime(5))
	fmt.Println(FindPrevPrime(28))
}

func FindPrevPrime(nb int) int {
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			nb -= 1
			FindPrevPrime(nb)
		}
	}
	return nb
}
