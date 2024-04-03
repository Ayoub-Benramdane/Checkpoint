package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args[1:]) != 1 {
		return
	}
	arg := os.Args[1]
	res := ""
	resfinal := ""
	for i := 0; i < len(arg); i++ {
		if arg[0] == 'a' || arg[0] == 'e' || arg[0] == 'i' || arg[0] == 'o' || arg[0] == 'u' || arg[0] == 'A' || arg[0] == 'E' || arg[0] == 'I' || arg[0] == 'O' || arg[0] == 'U' {
			resfinal += string(arg[i])
		} else if arg[i] != 'a' && arg[i] != 'e' && arg[i] != 'i' && arg[i] != 'o' && arg[i] != 'u' && arg[i] != 'A' && arg[i] != 'E' && arg[i] != 'I' && arg[i] != 'O' && arg[i] != 'U' && resfinal == "" {
			res += string(arg[i])
		} else {
			resfinal += string(arg[i])
		}
	}
	if resfinal == "" {
		fmt.Println("No vowels")
		return
	}
	resfinal += res + "ay"
	fmt.Println(resfinal)
}
