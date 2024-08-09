package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	vowels := "aeiou"
	str := ""
	strFinal := ""
	for i, c := range os.Args[1] {
		for _, v := range vowels {
			if c == v || c == v-32 {
				if str == "" {
					fmt.Println(os.Args[1] + "ay")
				} else {
					fmt.Println(os.Args[1][i:] + str + "ay")
				}
				return
			}
		}
		str += string(c)
	}
	if strFinal == "" {
		fmt.Println("No vowels")
		return
	}
}
