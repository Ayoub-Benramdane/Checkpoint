package main

import (
	"fmt"
	"os"
)

func main() {
	brackets := []rune{}
	for _, c := range os.Args[1:] {
		for _, v := range c {
			if v == '(' || v == '{' || v == '[' {
				brackets = append(brackets, v)
			} else if v == ')' {
				if brackets[len(brackets)-1] != '(' {
					fmt.Println("Error")
					return
				}
				brackets = brackets[:len(brackets)-1]
			} else if v == '}' {
				if brackets[len(brackets)-1] != '{' {
					fmt.Println("Error")
					return
				}
				brackets = brackets[:len(brackets)-1]
			} else if v == ']' {
				if brackets[len(brackets)-1] != '[' {
					fmt.Println("Error")
					return
				}
				brackets = brackets[:len(brackets)-1]
			}
		}
		fmt.Println("OK")
	}
}
