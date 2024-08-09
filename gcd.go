package main

import (
	"fmt"
)

func main() {
	fmt.Println(Gcd(42, 10))
	fmt.Println(Gcd(42, 12))
	fmt.Println(Gcd(14, 77))
	fmt.Println(Gcd(17, 3))
}

func Gcd(a, b uint) uint {
	var res uint
	n := 0
	if a > b {
		n = int(b)
	} else {
		n = int(a)
	}
	for i := 1; i <= n; i++ {
		if int(a)%i == 0 && int(b)%i == 0 {
			res = uint(i)
		}
	}
	return res
}
