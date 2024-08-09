package main

import (
	"fmt"
)

func main() {
	fmt.Print(FirstWord("hello there"))
	fmt.Print(FirstWord(""))
	fmt.Print(FirstWord("hello   .........  bye"))
}

func FirstWord(s string) string {
	str := ""
	for _, c := range s {
		if c == ' ' {
			return str + "\n"
		}
		str += string(c)
	}
	return str + "\n"
}
