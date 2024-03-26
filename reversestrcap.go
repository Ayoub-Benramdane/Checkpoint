package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	str := ""
	for _, c := range args {
		for _, v := range c {
			if v >= 'A' && v <= 'Z' {
				v += 32
			}
			str += string(v)
		}
		str += "\n"
	}
	myString := ""
	for i := 0; i < len(str); i++ {
		if i == len(str)-1 {
			break
		}
		if str[i+1] == ' ' {
			if str[i] >= 'a' && str[i] <= 'z' {
				myString += string(str[i] - 32)
			} else {
				myString += string(str[i])

			}
		} else {
			if str[i+1] == '\n' {
				if str[i] >= 'a' && str[i] <= 'z' {
					myString += string(str[i] - 32)
				} else {

					myString += string(str[i])
				}
			} else {
				myString += string(str[i])

			}
		}

	}
	fmt.Println(myString)
}
