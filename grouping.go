package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	expressions := []string{}
	exp := ""
	res := []string{}
	exist := false
	for i, c := range os.Args[1] {
		if i == 0 && c != '(' || i == len(os.Args[1])-1 && c != ')' {
			return
		}
		if c != '(' && c != ')' && c != '|' {
			exp += string(c)
		} else if exp != "" {
			expressions = append(expressions, exp)
			exp = ""
		}
	}
	for i, v := range os.Args[2] {
		if v >= 'a' && v <= 'z' || v >= 'A' && v <= 'Z' || v == '’' || v == '\'' {
			exp += string(v)
		}
		if exp != "" && (v == ' ' || i == len(os.Args[2])-1) {
			for _, c := range expressions {
				for i := 0; i <= len(exp)-len(c); i++ {
					if exp[i:i+len(c)] == c {
						exist = true
					}
					if exist {
						res = append(res, exp)
						exist = false
						break
					}
				}
			}
			exp = ""
		}
	}
	for i, c := range res {
		fmt.Printf("%d: %v", i+1, c)
		fmt.Println()
	}
}
