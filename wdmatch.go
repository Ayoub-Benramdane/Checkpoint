package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		return
	}
	indice := 0
	for _, c := range args[0] {
		valid := false
		for i := indice; i < len(args[1]); i++ {
			indice++
			if c == rune(args[1][i]) {
				valid = true
				break
			}
		}
		if !valid {
			return
		}
	}
	fmt.Println(args[0])
}
