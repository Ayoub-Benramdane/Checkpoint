package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		return
	}
	args1 := make(map[int]int)
	args2 := make(map[int]int)
	var res = []int{}
	resFinal := 0
	for i := 1; i <= ATOi(args[0]); i++ {
		if ATOi(args[0])%i == 0 {
			args1[int(i)]++
		}
	}
	for i := 1; i <= ATOi(args[1]); i++ {
		if ATOi(args[1])%i == 0 {
			args2[int(i)]++
		}
	}
	for c := range args1 {
		for v := range args2 {
			if c == v {
				res = append(res, c)
			}
		}
	}
	for i := 0; i < len(res); i++ {
		if resFinal < res[i] {
			resFinal = res[i]
		}
	}
	fmt.Println(resFinal)
}

func ITOa(nb int) string {
	res := ""
	for nb != 0 {
		res = string(nb%10+'0') + res
		nb /= 10
	}
	return res
}

func ATOi(s string) int {
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

