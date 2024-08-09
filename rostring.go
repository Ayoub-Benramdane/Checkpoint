package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	f_w := false
	var fWord, str, strFinal string
	for i, c := range os.Args[1] {
		if c == ' ' || i == len(os.Args[1])-1 {
			if c != ' ' {
				str += string(c)
			}
			if str != "" {
				if !f_w {
					fWord = str
					f_w = true
				} else {
					strFinal += str
					if check(os.Args[1], i+1) {
						strFinal += " "
					}
				}
				str = ""
			}
			if i == len(os.Args[1])-1 && fWord != "" && strFinal != "" {
				strFinal += " " + fWord
			} else if i == len(os.Args[1])-1 && fWord != "" {
				strFinal += fWord
			}
		} else {
			str += string(c)
		}
	}
	fmt.Println(strFinal)
}

func check(s string, indice int) bool {
	for i := indice; i < len(s); i++ {
		if s[i] > 32 && s[i] < 127 {
			return true
		}
	}
	return false
}
