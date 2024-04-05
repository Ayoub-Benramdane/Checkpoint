package main

import (
	"fmt"
)

func main() {
	s := "HelloHALhowHALareHALyouHAL?"
	fmt.Printf("%#v\n", Split(s, "HAL"))
}

func Split(s, sep string) []string {
	res := ""
	resFinal := []string{}
	for i := 0; i < len(s); i++ {
		if i+len(sep)+1 > len(s) {
			res += string(s[i])
		} else if s[i:i+len(sep)] != sep {
			res += string(s[i])
		} else {
			i += len(sep)-1
			if res != "" {
				resFinal = append(resFinal, res)
			}
			res = ""
		}
	}
	resFinal = append(resFinal, res)
	return resFinal
}
