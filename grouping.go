package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args[1:]) != 2 {
		return
	}
	arg1 := os.Args[1]
	arg2 := os.Args[2]
	res0 := ""
	res := []string{}
	res1 := ""
	resFinal := []string{}
	if arg1[0] != '(' || arg1[len(arg1)-1] != ')' {
		return
	}
	for _, c := range arg1 {
		if c != '|' && c != '(' && c != ')' {
			res0 += string(c)
		} else {
			if c != '(' {
				res = append(res, res0)
			}
			res0 = ""
		}
	}
	for i, c := range arg2 {
		if ((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '’' || c == '\'') && i < len(arg2)-1 {
			res1 += string(c)
		} else {
			if ((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '’' || c == '\'') && i == len(arg2)-1 {
				res1 += string(c)
			}
			for j := 0; j < len(res); j++ {
				for i := 0; i < len(res1)-len(res[j])+1; i++ {
					if res[j] == string(res1[i:i+len(res[j])]) {
						resFinal = append(resFinal, res1)
						break
					}
				}
			}
			res1 = ""
		}
	}
	for i := 0; i < len(resFinal); i++ {
		fmt.Printf("%d: %v", i+1, resFinal[i])
		fmt.Println()
	}
}
