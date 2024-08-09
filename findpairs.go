package main

import (
	"fmt"
	"os"
)

func main() {
	num := ""
	n := ""
	tab := []string{}
	tabFinal := []int{}
	numFinal := 0
	output := ""
	for i, c := range os.Args[2] {
		if i == 0 && c == '-' {
			num += string(c)
		} else if c < '0' || c > '9' {
			fmt.Println("Invalid target sum.")
			return
		} else {
			num += string(c)
		}
	}
	numFinal = atoi(num)
	for i, c := range os.Args[1] {
		if i == 0 && c != '[' || i == len(os.Args[1])-1 && c != ']' {
			fmt.Println("Invalid input.")
			return
		}
		if i != 0 {
			if (c == ',' || c == ']') && n != "" {
				tab = append(tab, n)
				n = ""
			} else if c != ',' && c != ' ' {
				if c >= '0' && c <= '9' || c == '-' {
					n += string(c)
				} else {
					n = string(c)
					fmt.Printf("Invalid number: %s", n)
					fmt.Println()
					return
				}
			}
		}
	}
	for _, c := range tab {
		tabFinal = append(tabFinal, atoi(c))
	}
	for i := 0; i < len(tabFinal); i++ {
		for j := i + 1; j < len(tabFinal); j++ {
			if tabFinal[i]+tabFinal[j] == numFinal {
				if output != "" {
					output += " "
				}
				output += "[" + itoa(i) + " " + itoa(j) + "]"
			}
		}
	}
	if output == "" {
		fmt.Printf("No pairs found.")
		fmt.Println()
		return
	}
	fmt.Printf("Pairs with sum %s: [%s]", num, output)
}

func atoi(s string) int {
	n := 0
	sign := 1
	for i, c := range s {
		if i == 0 && c == '-' {
			sign = -sign
		} else if c >= '0' && c <= '9' {
			n = n*10 + int(c-'0')
		}
	}
	return n * sign
}

func itoa(n int) string {
	res := ""
	if n == 0 {
		return "0"
	}
	for n != 0 {
		res = string(n%10+'0') + res
		n /= 10
	}
	return res
}
