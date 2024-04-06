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
	resFinal := "OK"
	for i := 0; i < len(args); i++ {
		for j := 0; j < len(args[i]); j++ {
			if args[i][j] == '(' || args[i][j] == '{' || args[i][j] == '[' || args[i][j] == ')' || args[i][j] == '}' || args[i][j] == ']' {
				res = append(res, string(args[i][j]))
			}
		}
		if res != nil {
			res1 = append(res1, res)
		} else {
			res1 = append(res1, []string{})
		}
		res = []string{}
	}
	fmt.Println(res1)
	for i := 0; i < len(res1); i++ {
		if len(res1[i]) == 0 {
			fmt.Println("OK1")
		} else {
			for i := 0; i < len(res); i++ {
				if res[i] == ")" {
					if res[i-1] == "(" {
						res = append(res[:i-1], res[i+1:]...)
						i -= 2
					} else {
						resFinal = "Error"
						break
					}
				} else if res[i] == "}" {
					if res[i-1] == "{" {
						res = append(res[:i-1], res[i+1:]...)
						i -= 2
					} else {
						resFinal = "Error"
						break
					}
				} else if res[i] == "]" {
					if res[i-1] == "[" {
						res = append(res[:i-1], res[i+1:]...)
						i -= 2
					} else {
						resFinal = "Error"
						break
					}
				}
			}
		}
	}
	fmt.Println(resFinal)
}
