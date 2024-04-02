package main

import "fmt"

func main() {
	fmt.Println(FindPrevPrime(5))
	fmt.Println(FindPrevPrime(4))
}

func FindPrevPrime(nb int) int {
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			nb = nb - 1
			FindPrevPrime(nb - 1)
		}
	}
	return nb
}
