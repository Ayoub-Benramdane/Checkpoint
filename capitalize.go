package main

import (
	"fmt"
)

func main() {
	fmt.Println(Capitalize("Hello! How are you? How+are+things+4you?"))
}

func Capitalize(s string) string {
	isFirst := true
	res := ""
	for _, c := range s {
		if isFirst {
			if c >= 'a' && c <= 'z' {
				res += string(c - 32)
			} else {
				res += string(c)
			}
			isFirst = false
		} else if c >= 'A' && c <= 'Z' {
			res += string(c + 32)
		} else {
			res += string(c)
		}
		if c < '0' || c > '9' && c < 'A' || c > 'Z' && c < 'a' || c > 'z' {
			isFirst = true
		}
	}
	return res
}
