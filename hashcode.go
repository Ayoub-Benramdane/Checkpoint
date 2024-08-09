package main

import (
	"fmt"
)

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}

func HashCode(dec string) string {
	hash := ""
	length := len(dec)
	for i := 0; i < length; i++ {
		hashValue := (int(dec[i]) + length) % 127
		if hashValue < 32 {
			hashValue += 33
		}
		hash += string(rune(hashValue))
	}
	return hash
}
