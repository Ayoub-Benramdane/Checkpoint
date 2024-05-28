package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args[1]
	bracketsOp := []rune{'(', '[', '{'}
	bracketsClo := []rune{')', ']', '}'}
	index := []int{}
	tab := []rune{}
	for _, c := range arg {
		for i, b := range bracketsOp {
			if b == c {
				tab = append(tab, b)
				index = append(index, i)
			}
			continue
		}
		for i, v := range bracketsClo {
			if v == c {
				if len(tab) == 0 {
					fmt.Println("Error")
					os.Exit(0)
				}
				if len(tab) > 0 && i == index[len(index)-1] {
					tab = tab[:len(tab)-1]
					index = index[:len(index)-1]
				} else {
					fmt.Println("Error")
					os.Exit(0)
				}
			}
		}
	}
	if len(tab) > 0 {
		fmt.Println("Error")
	} else {
		fmt.Println("OK")
	}
}

