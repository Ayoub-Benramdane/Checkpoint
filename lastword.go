package main

import (
	"fmt"
)

func main() {
	fmt.Print(LastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Print(LastWord(" lorem,ipsum "))
	fmt.Print(LastWord(" "))
}

func LastWord(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' && check(i, s) {
			res = ""
			continue
		}
		if s[i] != ' ' {
			res += string(s[i])
		}
	}
	return res + "\n"
}

func check(indice int, s string) bool {
	for i := indice; i < len(s); i++ {
		if s[i] != ' ' {
			return true
		}
	}
	return false
}
