package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		return
	}
	res := []string{}
	res1 := [][]string{}
	resFinal := []string{}
	for i := 0; i < len(args); i++ {
		for j := 0; j < len(args[i]); j++ {
			if args[i][j] == '(' || args[i][j] == '{' || args[i][j] == '[' || args[i][j] == ')' || args[i][j] == '}' || args[i][j] == ']' {
				res = append(res, string(args[i][j]))
			}
		}
		if len(args[i]) == 0 {
			res1 = append(res1, nil)
		} else {
			res1 = append(res1, res)
		}
		res = []string{}
	}

	for i := 0; i < len(res1); i++ {
		if len(res1[i])%2 != 0 {
			resFinal = append(resFinal, "Error")
			continue
		}
		for i := 0; i < len(res); i++ {
			if res[i] == ")" {
				if res[i-1] == "(" {
					res = append(res[:i-1], res[i+1:]...)
					i -= 2
				} else {
					resFinal = append(resFinal, "Error")
					break
				}
			} else if res[i] == "}" {
				if res[i-1] == "{" {
					res = append(res[:i-1], res[i+1:]...)
					i -= 2
				} else {
					resFinal = append(resFinal, "Error")
					break
				}
			} else if res[i] == "]" {
				if res[i-1] == "[" {
					res = append(res[:i-1], res[i+1:]...)
					i -= 2
				} else {
					resFinal = append(resFinal, "Error")
					break
				}
			}
		}
		resFinal = append(resFinal, "OK")
	}
	for i := 0; i < len(resFinal); i++ {
		fmt.Println(resFinal[i])
	}
}
