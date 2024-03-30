package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	for i := 1; i <= 9; i++ {
		fmt.Println(i, "x", args[0], "=", IToa(i*AToi(args[0])))
	}
}

func IToa(nb int) string {
	res := ""
	for nb != 0 {
		res = string(nb%10+'0') + res
		nb /= 10
	}
	return res
}

func AToi(s string) int {
	t := 1
	res := 0
	for i, c := range s {
		if c == '-' && i == 0 {
			t = -1
			continue
		} else if c == '+' && i == 0 {
			continue
		} else if c >= '0' && c <= '9' {
			res = res*10 + int(c-'0')
		} else {
			return 0
		}
	}
	return res * t
}
