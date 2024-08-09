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
	str := ""
	for i, c := range s {
		if c == ' ' && checkAlpha(s, i) {
			str = ""
		} else {
			str += string(c)
		}
	}
	return str + "\n"
}

func checkAlpha(s string, indice int) bool {
	for i := indice; i < len(s); i++ {
		if s[i] > 32 && s[i] < 127 {
			return true
		}
	}
	return false
}
