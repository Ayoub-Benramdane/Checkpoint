package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	res := []string{}
	str := ""
	for i, c := range os.Args[1] {
		if c == ' ' || i == len(os.Args[1])-1 {
			if c != ' ' {
				str += string(c)
			}
			if str != "" {
				res = append(res, str)
				str = ""
			}
		} else {
			str += string(c)
		}
	}
	for i := len(res) - 1; i >= 0; i-- {
		str += res[i]
		if i != 0 {
			str += " "
		}
	}
	fmt.Println(str)
}
