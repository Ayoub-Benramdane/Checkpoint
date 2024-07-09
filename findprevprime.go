package main

import "fmt"

func main() {
	//fmt.Println(FindPrevPrime(225225))
	fmt.Println(FindPrevPrime(28))
}

var c = 0

func FindPrevPrime(nb int) int {
	c++
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			fmt.Println(c, nb)
			nb -= 1
			i = 2
			FindPrevPrime(nb)
		}
		if nb%i != 0 && i+1 == nb {
			return nb
		}

	}
	return nb
}
